/*
 * Copyright (c) 2019  InterDigital Communications, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package netchar

import (
	"errors"
	"strconv"
	"sync"
	"time"

	dkm "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-key-mgr"
	dataModel "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-model"
	log "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger"
	mod "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-model"
	mq "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-mq"
	pss "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-pdu-session-store"
	redis "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-redis"
)

const netCharControlDb = 0
const defaultTickerPeriod int = 500
const NetCharControls string = "net-char-controls"
const NetCharControlChannel string = NetCharControls

// Callback function types
type NetCharUpdateCb func(string, string, float64, float64, float64, string, float64)
type UpdateCompleteCb func()

// NetChar Interface
type NetCharMgr interface {
	Register(NetCharUpdateCb, UpdateCompleteCb)
	Start() error
	Stop()
	IsRunning() bool
}

// NetCharAlgo
type NetCharAlgo interface {
	ProcessScenario(*mod.Model, map[string]map[string]*dataModel.PduSessionInfo) error
	CalculateNetChar() []FlowNetChar
	SetConfigAttribute(string, string)
}

// NetChar
type NetChar struct {
	Latency      float64
	Jitter       float64
	PacketLoss   float64
	Throughput   float64
	Distribution string
}

// NetChar
type ElemNetChar struct {
	Latency      float64
	Jitter       float64
	Distribution string
	PacketLoss   float64
	ThroughputUl float64
	ThroughputDl float64
}

// FlowNetChar
type FlowNetChar struct {
	SrcElemName string
	DstElemName string
	MyNetChar   NetChar
}

// NetCharConfig
type NetCharConfig struct {
	Action              string
	RecalculationPeriod int
	LogVerbose          bool
}

// NetCharManager Object
type NetCharManager struct {
	name             string
	namespace        string
	baseKey          string
	isStarted        bool
	ticker           *time.Ticker
	rc               *redis.Connector
	pduSessionStore  *pss.PduSessionStore
	pduSessions      map[string]map[string]*dataModel.PduSessionInfo
	mutex            sync.Mutex
	config           NetCharConfig
	mqLocal          *mq.MsgQueue
	activeModel      *mod.Model
	netCharUpdateCb  NetCharUpdateCb
	updateCompleteCb UpdateCompleteCb
	algo             NetCharAlgo
	handlerId        int
}

// NewNetChar - Create, Initialize and connect
func NewNetChar(name string, namespace string, redisAddr string) (*NetCharManager, error) {
	// Create new instance & set default config
	var err error
	var ncm NetCharManager
	if name == "" {
		err = errors.New("Missing name")
		log.Error(err)
		return nil, err
	}
	ncm.name = name
	ncm.namespace = namespace
	ncm.baseKey = dkm.GetKeyRoot(namespace)
	ncm.isStarted = false
	ncm.config.RecalculationPeriod = defaultTickerPeriod

	// Create message queue
	ncm.mqLocal, err = mq.NewMsgQueue(mq.GetLocalName(namespace), name, namespace, redisAddr)
	if err != nil {
		log.Error("Failed to create Message Queue with error: ", err)
		return nil, err
	}
	log.Info("Message Queue created")

	// Create new NetCharAlgo
	ncm.algo, err = NewSegmentAlgorithm(ncm.name, ncm.namespace, redisAddr)
	if err != nil {
		log.Error("Failed to create NetCharAlgo with error: ", err)
		return nil, err
	}

	// Create new Model
	modelCfg := mod.ModelCfg{
		Name:      "activeScenario",
		Namespace: ncm.namespace,
		Module:    name,
		UpdateCb:  nil,
		DbAddr:    redisAddr,
	}
	ncm.activeModel, err = mod.NewModel(modelCfg)
	if err != nil {
		log.Error("Failed to create model: ", err.Error())
		return nil, err
	}

	// Create new Control listener
	ncm.rc, err = redis.NewConnector(redisAddr, netCharControlDb)
	if err != nil {
		log.Error("Failed connection to redis DB. Error: ", err)
		return nil, err
	}
	log.Info("Connected to Control Listener redis DB")

	// Connect to PDU Session Store
	ncm.pduSessionStore, err = pss.NewPduSessionStore(ncm.namespace, redisAddr)
	if err != nil {
		log.Error("Failed connection to PDU Session Store: ", err.Error())
		return nil, err
	}
	log.Info("Connected to PDU Session Store")

	// Register Message Queue handler
	handler := mq.MsgHandler{Handler: ncm.msgHandler, UserData: nil}
	ncm.handlerId, err = ncm.mqLocal.RegisterHandler(handler)
	if err != nil {
		log.Error("Failed to listen for sandbox updates: ", err.Error())
		return nil, err
	}

	// Listen for Control updates
	err = ncm.rc.Subscribe(NetCharControlChannel)
	if err != nil {
		log.Error("Failed to subscribe to Pub/Sub events on NetCharControlChannel. Error: ", err)
		return nil, err
	}
	go func() {
		_ = ncm.rc.Listen(ncm.eventHandler)
	}()

	log.Debug("NetChar successfully created: ", ncm.name)
	return &ncm, nil
}

// Register - Register NetChar callback functions
func (ncm *NetCharManager) Register(netCharUpdateCb NetCharUpdateCb, updateCompleteCb UpdateCompleteCb) {
	ncm.netCharUpdateCb = netCharUpdateCb
	ncm.updateCompleteCb = updateCompleteCb
}

// Start - Start NetChar
func (ncm *NetCharManager) Start() error {
	if !ncm.isStarted {
		ncm.isStarted = true

		// Process current controls
		ncm.updateControls()

		// Process current pdu sessions
		go ncm.processPduSessionUpdate()

		// Process current scenario
		go ncm.processActiveScenarioUpdate()

		// Start ticker to refresh net char periodically
		ncm.ticker = time.NewTicker(time.Duration(ncm.config.RecalculationPeriod) * time.Millisecond)
		go func() {
			for range ncm.ticker.C {
				ncm.mutex.Lock()
				if ncm.isStarted {
					ncm.updateNetChars()
				}
				ncm.mutex.Unlock()
			}
		}()
		log.Debug("Network Characteristics Manager started: ", ncm.name)
	}
	return nil
}

// Stop - Stop NetChar
func (ncm *NetCharManager) Stop() {
	if ncm.isStarted {
		ncm.isStarted = false
		ncm.ticker.Stop()
		log.Debug("NetChar stopped ", ncm.name)
	}
}

// IsRunning
func (ncm *NetCharManager) IsRunning() bool {
	return ncm.isStarted
}

// Message Queue handler
func (ncm *NetCharManager) msgHandler(msg *mq.Msg, userData interface{}) {
	switch msg.Message {
	case mq.MsgScenarioActivate:
		log.Debug("RX MSG: ", mq.PrintMsg(msg))
		ncm.processActiveScenarioUpdate()
	case mq.MsgScenarioUpdate:
		log.Debug("RX MSG: ", mq.PrintMsg(msg))
		ncm.processActiveScenarioUpdate()
	case mq.MsgScenarioTerminate:
		log.Debug("RX MSG: ", mq.PrintMsg(msg))
		ncm.processActiveScenarioUpdate()
	case mq.MsgPduSessionCreated, mq.MsgPduSessionTerminated:
		log.Debug("RX MSG: ", mq.PrintMsg(msg))
		ncm.processPduSessionUpdate()
	default:
		log.Trace("Ignoring unsupported message: ", mq.PrintMsg(msg))
	}
}

// eventHandler - Events received and processed by the registered channels
func (ncm *NetCharManager) eventHandler(channel string, payload string) {
	// Handle Message according to Rx Channel
	switch channel {
	case NetCharControlChannel:
		log.Debug("Event received on channel: ", channel, " payload: ", payload)
		ncm.updateControls()
	default:
		log.Warn("Unsupported channel")
	}
}

// processActiveScenarioUpdate
func (ncm *NetCharManager) processActiveScenarioUpdate() {
	// Sync with active scenario store
	ncm.activeModel.UpdateScenario()

	ncm.mutex.Lock()
	defer ncm.mutex.Unlock()

	if ncm.isStarted {
		// Process updated scenario using algorithm
		err := ncm.algo.ProcessScenario(ncm.activeModel, ncm.pduSessions)
		if err != nil {
			log.Error("Failed to process active model with error: ", err)
			return
		}

		// Recalculate network characteristics
		ncm.updateNetChars()
	}
}

// processPduSessionUpdate
func (ncm *NetCharManager) processPduSessionUpdate() {

	ncm.mutex.Lock()
	defer ncm.mutex.Unlock()

	// Refresh PDU session cache
	var err error
	ncm.pduSessions, err = ncm.pduSessionStore.GetAllPduSessions()
	if err != nil {
		log.Error("Failed to retrieve PDU session maps with error: ", err)
		return
	}

	if ncm.isStarted {
		// Process updated scenario using algorithm
		err := ncm.algo.ProcessScenario(ncm.activeModel, ncm.pduSessions)
		if err != nil {
			log.Error("Failed to process active model with error: ", err)
			return
		}

		// Recalculate network characteristics
		ncm.updateNetChars()
	}
}

// updateNetChars
func (ncm *NetCharManager) updateNetChars() {
	// Recalculate network characteristics
	updatedNetCharList := ncm.algo.CalculateNetChar()

	// Apply updates, if any
	if len(updatedNetCharList) != 0 {
		for _, flowNetChar := range updatedNetCharList {
			if ncm.netCharUpdateCb != nil {
				ncm.netCharUpdateCb(flowNetChar.DstElemName, flowNetChar.SrcElemName, flowNetChar.MyNetChar.Throughput, flowNetChar.MyNetChar.Latency, flowNetChar.MyNetChar.Jitter, flowNetChar.MyNetChar.Distribution /*flowNetChar.MyNetChar.Distribution,*/, flowNetChar.MyNetChar.PacketLoss)
			}
		}
		if ncm.updateCompleteCb != nil {
			ncm.updateCompleteCb()
		}
	}
}

// updateControls - Update all the different configurations attributes based on the content of the DB for dynamic updates
func (ncm *NetCharManager) updateControls() {
	ncm.mutex.Lock()
	var controls = make(map[string]interface{})
	keyName := ncm.baseKey + NetCharControls
	err := ncm.rc.ForEachEntry(keyName, ncm.getControlsEntryHandler, controls)
	if err != nil {
		log.Error("Failed to get entries: ", err)
	}
	ncm.mutex.Unlock()
}

// getControlsEntryHandler - Update all the different configurations attributes based on the content of the DB for dynamic updates
func (ncm *NetCharManager) getControlsEntryHandler(key string, fields map[string]string, userData interface{}) (err error) {

	actionName := ""
	tickerPeriod := defaultTickerPeriod
	logVerbose := false

	for fieldName, fieldValue := range fields {
		switch fieldName {
		case "action":
			actionName = fieldValue
		case "recalculationPeriod":
			tickerPeriod, err = strconv.Atoi(fieldValue)
			if err != nil {
				tickerPeriod = defaultTickerPeriod
			}
		case "logVerbose":
			if "yes" == fieldValue {
				logVerbose = true
			}
		default:
		}
		ncm.algo.SetConfigAttribute(fieldName, fieldValue)
	}

	ncm.config.Action = actionName
	ncm.config.RecalculationPeriod = tickerPeriod
	ncm.config.LogVerbose = logVerbose

	ncm.applyAction()
	return nil
}

// applyAction - Execute the action in the configuration parameters for controls on the NetChar object
func (ncm *NetCharManager) applyAction() {
	switch ncm.config.Action {
	case "start":
		if !ncm.isStarted {
			_ = ncm.Start()
		}
	case "stop":
		if ncm.isStarted {
			ncm.Stop()
		}
	default:
	}
}
