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
	"fmt"
	"testing"

	dkm "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-key-mgr"
	dataModel "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-model"
	log "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger"
	mod "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-model"
	redis "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-redis"
)

const segAlgoRedisAddr string = "localhost:30380"
const testModuleName string = "test-net-char-mgr"
const testModuleNamespace string = "test-ns"

var jsonTestScenario = `{"version":"1.5.3","name":"demo1","deployment":{"netChar":{"latency":50,"latencyVariation":5,"latencyDistribution":"Normal","throughputDl":1000,"throughputUl":1000},"domains":[{"id":"PUBLIC","name":"PUBLIC","type":"PUBLIC","netChar":{"latency":6,"latencyVariation":2,"throughputDl":1000000,"throughputUl":1000000},"zones":[{"id":"PUBLIC-COMMON","name":"PUBLIC-COMMON","type":"COMMON","netChar":{"latency":5,"latencyVariation":1},"networkLocations":[{"id":"PUBLIC-COMMON-DEFAULT","name":"PUBLIC-COMMON-DEFAULT","type":"DEFAULT","netChar":{"latency":1,"latencyVariation":1,"throughputDl":50000,"throughputUl":50000,"packetLoss":1},"physicalLocations":[{"id":"cloud","name":"cloud","type":"DC","connected":true,"processes":[{"id":"cloud-iperf","name":"cloud-iperf","type":"CLOUD-APP","image":"meep-docker-registry:30001/iperf-server","commandArguments":"-c, export; iperf -s -p $IPERF_SERVICE_PORT","commandExe":"/bin/bash","serviceConfig":{"name":"cloud-iperf","ports":[{"protocol":"UDP","port":80}]},"netChar":{}},{"id":"cloud-svc","name":"cloud-svc","type":"CLOUD-APP","image":"meep-docker-registry:30001/demo-server","environment":"MGM_GROUP_NAME=cloud-svc, MGM_APP_ID=cloud-svc, MGM_APP_PORT=80","serviceConfig":{"name":"cloud-svc","ports":[{"protocol":"TCP","port":80}]},"netChar":{}}],"netChar":{}}]}]}]},{"id":"operator1","name":"operator1","type":"OPERATOR","netChar":{"latency":15,"latencyVariation":3,"throughputDl":1000,"throughputUl":1000},"zones":[{"id":"operator1-COMMON","name":"operator1-COMMON","type":"COMMON","netChar":{"latency":5,"latencyVariation":1},"networkLocations":[{"id":"operator1-COMMON-DEFAULT","name":"operator1-COMMON-DEFAULT","type":"DEFAULT","netChar":{"latency":1,"latencyVariation":1,"throughputDl":50000,"throughputUl":50000,"packetLoss":1}}]},{"id":"zone1","name":"zone1","type":"ZONE","netChar":{"latency":5,"latencyVariation":1},"networkLocations":[{"id":"zone1-DEFAULT","name":"zone1-DEFAULT","type":"DEFAULT","netChar":{"latency":1,"latencyVariation":1,"throughputDl":50000,"throughputUl":50000,"packetLoss":1},"physicalLocations":[{"id":"zone1-edge1","name":"zone1-edge1","type":"EDGE","connected":true,"processes":[{"id":"zone1-edge1-iperf","name":"zone1-edge1-iperf","type":"EDGE-APP","image":"meep-docker-registry:30001/iperf-server","commandArguments":"-c, export; iperf -s -p $IPERF_SERVICE_PORT","commandExe":"/bin/bash","serviceConfig":{"name":"zone1-edge1-iperf","meSvcName":"iperf","ports":[{"protocol":"UDP","port":80}]},"netChar":{}},{"id":"zone1-edge1-svc","name":"zone1-edge1-svc","type":"EDGE-APP","image":"meep-docker-registry:30001/demo-server","environment":"MGM_GROUP_NAME=svc, MGM_APP_ID=zone1-edge1-svc, MGM_APP_PORT=80","serviceConfig":{"name":"zone1-edge1-svc","meSvcName":"svc","ports":[{"protocol":"TCP","port":80}]},"netChar":{}}],"netChar":{}}]},{"id":"zone1-poa1","name":"zone1-poa1","type":"POA","netChar":{"latency":1,"latencyVariation":1,"throughputDl":1000,"throughputUl":1000},"physicalLocations":[{"id":"zone1-fog1","name":"zone1-fog1","type":"FOG","connected":true,"processes":[{"id":"zone1-fog1-iperf","name":"zone1-fog1-iperf","type":"EDGE-APP","image":"meep-docker-registry:30001/iperf-server","commandArguments":"-c, export; iperf -s -p $IPERF_SERVICE_PORT;","commandExe":"/bin/bash","serviceConfig":{"name":"zone1-fog1-iperf","meSvcName":"iperf","ports":[{"protocol":"UDP","port":80}]},"netChar":{}},{"id":"zone1-fog1-svc","name":"zone1-fog1-svc","type":"EDGE-APP","image":"meep-docker-registry:30001/demo-server","environment":"MGM_GROUP_NAME=svc, MGM_APP_ID=zone1-fog1-svc, MGM_APP_PORT=80","serviceConfig":{"name":"zone1-fog1-svc","meSvcName":"svc","ports":[{"protocol":"TCP","port":80}]},"netChar":{}}],"netChar":{}},{"id":"ue1","name":"ue1","type":"UE","connected":true,"wireless":true,"wirelessType":"wifi,5g,4g","processes":[{"id":"ue1-iperf","name":"ue1-iperf","type":"UE-APP","image":"meep-docker-registry:30001/iperf-client","commandArguments":"-c, export; iperf -u -c $IPERF_SERVICE_HOST -p $IPERF_SERVICE_PORT -t 3600 -b 50M;","commandExe":"/bin/bash","netChar":{}}],"netChar":{}},{"id":"ue2-ext","name":"ue2-ext","type":"UE","isExternal":true,"connected":true,"wireless":true,"wirelessType":"wifi,5g,4g","processes":[{"id":"ue2-svc","name":"ue2-svc","type":"UE-APP","isExternal":true,"externalConfig":{"ingressServiceMap":[{"name":"svc","port":80,"externalPort":31111,"protocol":"TCP"},{"name":"iperf","port":80,"externalPort":31222,"protocol":"UDP"},{"name":"cloud-svc","port":80,"externalPort":31112,"protocol":"TCP"},{"name":"cloud-iperf","port":80,"externalPort":31223,"protocol":"UDP"}]},"netChar":{}}],"netChar":{}}]},{"id":"zone1-poa2","name":"zone1-poa2","type":"POA","netChar":{"latency":10,"latencyVariation":2,"throughputDl":50,"throughputUl":50}}]},{"id":"zone2","name":"zone2","type":"ZONE","netChar":{"latency":5,"latencyVariation":1},"networkLocations":[{"id":"zone2-DEFAULT","name":"zone2-DEFAULT","type":"DEFAULT","netChar":{"latency":1,"latencyVariation":1,"throughputDl":50000,"throughputUl":50000,"packetLoss":1},"physicalLocations":[{"id":"zone2-edge1","name":"zone2-edge1","type":"EDGE","connected":true,"processes":[{"id":"zone2-edge1-iperf","name":"zone2-edge1-iperf","type":"EDGE-APP","image":"meep-docker-registry:30001/iperf-server","commandArguments":"-c, export; iperf -s -p $IPERF_SERVICE_PORT;","commandExe":"/bin/bash","serviceConfig":{"name":"zone2-edge1-iperf","meSvcName":"iperf","ports":[{"protocol":"UDP","port":80}]},"netChar":{}},{"id":"zone2-edge1-svc","name":"zone2-edge1-svc","type":"EDGE-APP","image":"meep-docker-registry:30001/demo-server","environment":"MGM_GROUP_NAME=svc, MGM_APP_ID=zone2-edge1-svc, MGM_APP_PORT=80","serviceConfig":{"name":"zone2-edge1-svc","meSvcName":"svc","ports":[{"protocol":"TCP","port":80}]},"netChar":{}}],"netChar":{}}]},{"id":"zone2-poa1","name":"zone2-poa1","type":"POA","netChar":{"latency":1,"latencyVariation":1,"throughputDl":20,"throughputUl":20}}]}]}]}}`

func TestSegAlgoSegmentation(t *testing.T) {
	fmt.Println("--- ", t.Name())
	log.MeepTextLogInit(t.Name())

	// Create Model & add Scenario to use for testing
	fmt.Println("Create Model")
	modelCfg := mod.ModelCfg{Name: "activeScenario", Namespace: testModuleNamespace, Module: testModuleName, UpdateCb: nil, DbAddr: segAlgoRedisAddr}
	activeModel, err := mod.NewModel(modelCfg)
	if err != nil {
		t.Fatalf("Failed to create Model instance")
	}
	fmt.Println("Set scenario in Model")
	err = activeModel.SetScenario([]byte(jsonTestScenario))
	if err != nil {
		t.Fatalf("Failed to set scenario in model")
	}

	// Create new Algorithm
	fmt.Println("Create new algorithm")
	algo, err := NewSegmentAlgorithm(testModuleName, testModuleNamespace, segAlgoRedisAddr)
	if err != nil {
		t.Fatalf("Failed to create a SegAlgo object.")
	}
	if len(algo.FlowMap) != 0 {
		t.Fatalf("Flow Map not empty")
	}
	if len(algo.SegmentMap) != 0 {
		t.Fatalf("Segment Map not empty")
	}

	// Test Algorithm
	fmt.Println("Test algo without scenario")
	updatedNetCharList := algo.CalculateNetChar()
	if len(updatedNetCharList) != 0 {
		t.Fatalf("Updated net char list not empty")
	}

	fmt.Println("Process scenario model")
	err = algo.ProcessScenario(activeModel)
	if err != nil {
		t.Fatalf("Failed to process scenario model")
	}
	if len(algo.FlowMap) != 90 {
		t.Fatalf("Invalid Flow Map entry count")
	}

	if len(algo.SegmentMap) != 42 {
		t.Fatalf("Invalid Segment Map entry count")
	}

	// Validate algorithm segmentation
	fmt.Println("Validate algorithm segmentation")
	if !validatePath(algo, "zone1-fog1-iperf:ue1-iperf", 3) {
		t.Fatalf("Invalid path")
	}
	if !validatePath(algo, "zone2-edge1-iperf:ue1-iperf", 7) {
		t.Fatalf("Invalid path")
	}

	// Validate algorithm Calculations
	fmt.Println("Test algo calculation with some flows updated with metrics")
	updatedNetCharList = algo.CalculateNetChar()
	if len(updatedNetCharList) != 90 {
		t.Fatalf("Updated net char list not fully filled")
	}

	fmt.Println("Test algo calculation without changes in metrics")
	updatedNetCharList = algo.CalculateNetChar()
	if len(updatedNetCharList) != 0 {
		t.Fatalf("Updated net char list not empty")
	}

	// Verify algo after Network Characteristic update
	fmt.Println("Update Net char")
	var netCharUpdateEvent dataModel.EventNetworkCharacteristicsUpdate
	netCharUpdateEvent.ElementName = "zone1-poa1"
	netCharUpdateEvent.ElementType = "POA"
	var netChar dataModel.NetworkCharacteristics
	netChar.ThroughputUl = 100
	netCharUpdateEvent.NetChar = &netChar
	err = activeModel.UpdateNetChar(&netCharUpdateEvent)
	if err != nil {
		t.Fatalf("Error updating net char")
	}
	err = algo.ProcessScenario(activeModel)
	if err != nil {
		t.Fatalf("Failed to process scenario model")
	}
	if len(algo.FlowMap) != 90 {
		t.Fatalf("Invalid Flow Map entry count")
	}
	if len(algo.SegmentMap) != 42 {
		t.Fatalf("Invalid Segment Map entry count")
	}
	updatedNetCharList = algo.CalculateNetChar()
	if len(updatedNetCharList) != 34 {
		t.Fatalf("Updated net char list not empty")
	}

	// Verify algo after UE Mobility event update
	fmt.Println("Move ue1")
	_, _, err = activeModel.MoveNode("ue1", "zone2-poa1")
	if err != nil {
		t.Fatalf("Error moving UE")
	}
	err = algo.ProcessScenario(activeModel)
	if err != nil {
		t.Fatalf("Failed to process scenario model")
	}
	if len(algo.FlowMap) != 90 {
		t.Fatalf("Invalid Flow Map entry count")
	}

	if len(algo.SegmentMap) != 44 {
		t.Fatalf("Invalid Segment Map entry count")
	}

	// Validate algorithm segmentation
	fmt.Println("Validate algorithm segmentation")
	if !validatePath(algo, "zone1-fog1-iperf:ue1-iperf", 7) {
		t.Fatalf("Invalid path")
	}
	if !validatePath(algo, "zone2-edge1-iperf:ue1-iperf", 5) {
		t.Fatalf("Invalid path")
	}

	// Validate algorithm Calculations
	updatedNetCharList = algo.CalculateNetChar()
	if len(updatedNetCharList) != 18 {
		t.Fatalf("Updated net char list not empty")
	}

	// Verify algo after model update
	fmt.Println("Move ue1")
	_, _, err = activeModel.MoveNode("ue1", "zone1-poa1")
	if err != nil {
		t.Fatalf("Error moving UE")
	}
	err = algo.ProcessScenario(activeModel)
	if err != nil {
		t.Fatalf("Failed to process scenario model")
	}
	if len(algo.FlowMap) != 90 {
		t.Fatalf("Invalid Flow Map entry count")
	}
	if len(algo.SegmentMap) != 42 {
		t.Fatalf("Invalid Segment Map entry count")
	}

	// Validate algorithm segmentation
	fmt.Println("Validate algorithm segmentation")
	if !validatePath(algo, "zone1-fog1-iperf:ue1-iperf", 3) {
		t.Fatalf("Invalid path")
	}
	if !validatePath(algo, "zone2-edge1-iperf:ue1-iperf", 7) {
		t.Fatalf("Invalid path")
	}

	// Validate algorithm Calculations
	updatedNetCharList = algo.CalculateNetChar()
	if len(updatedNetCharList) != 18 {
		t.Fatalf("Updated net char list not empty")
	}

	// Clear model and make sure all
	fmt.Println("Create new Model instance")
	modelCfg = mod.ModelCfg{Name: "activeScenario", Namespace: testModuleNamespace, Module: testModuleName, UpdateCb: nil, DbAddr: segAlgoRedisAddr}
	activeModel, err = mod.NewModel(modelCfg)
	if err != nil {
		t.Fatalf("Failed to create Model instance")
	}
	fmt.Println("Process empty scenario model")
	err = algo.ProcessScenario(activeModel)
	if err != nil {
		t.Fatalf("Failed to process scenario model")
	}
	if len(algo.FlowMap) != 0 {
		t.Fatalf("Invalid Flow Map entry count")
	}
	if len(algo.SegmentMap) != 0 {
		t.Fatalf("Invalid Segment Map entry count")
	}
	fmt.Println("Test algo without scenario")
	updatedNetCharList = algo.CalculateNetChar()
	if len(updatedNetCharList) != 0 {
		t.Fatalf("Updated net char list not empty")
	}
}

func TestSegAlgoCalculation(t *testing.T) {
	fmt.Println("--- ", t.Name())
	log.MeepTextLogInit(t.Name())

	// Create connection to Metrics Redis DB to inject metrics
	rc, err := redis.NewConnector(segAlgoRedisAddr, metricsDb)
	if err != nil {
		t.Fatalf("Failed connection to Metrics redis DB")
	}

	// Create Model & add Scenario to use for testing
	fmt.Println("Create Model")
	modelCfg := mod.ModelCfg{Name: "activeScenario", Namespace: testModuleNamespace, Module: testModuleName, UpdateCb: nil, DbAddr: segAlgoRedisAddr}
	activeModel, err := mod.NewModel(modelCfg)
	if err != nil {
		t.Fatalf("Failed to create Model instance")
	}
	fmt.Println("Set scenario in Model")
	err = activeModel.SetScenario([]byte(jsonTestScenario))
	if err != nil {
		t.Fatalf("Failed to set scenario in model")
	}

	// Create & Process new Algorithm
	fmt.Println("Create new algorithm")
	algo, err := NewSegmentAlgorithm(testModuleName, testModuleNamespace, segAlgoRedisAddr)
	if err != nil {
		t.Fatalf("Failed to create a SegAlgo object.")
	}
	fmt.Println("Process scenario model")
	err = algo.ProcessScenario(activeModel)
	if err != nil {
		t.Fatalf("Failed to process scenario model")
	}

	// Validate algorithm Calculations
	fmt.Println("Test algorithm calculations with & without metrics")
	updatedNetCharList := algo.CalculateNetChar()
	if len(updatedNetCharList) != 90 {
		t.Fatalf("Updated net char list not fully filled")
	}

	if !validateNetCharUpdate(updatedNetCharList, "cloud-iperf", "ue1-iperf", 121, 15, 0, 200) {
		t.Fatalf("Error in Net Char initial calculation")
	}

	// Update metrics & recalculate
	if !setMetrics(rc, "zone1-fog1-iperf", "ue1-iperf", 100) {
		t.Fatalf("Error updating metrics")
	}
	if !setMetrics(rc, "zone1-fog1-svc", "ue2-svc", 100) {
		t.Fatalf("Error updating metrics")
	}
	updatedNetCharList = algo.CalculateNetChar()
	if len(updatedNetCharList) != 2 {
		t.Fatalf("Invalid net char update list")
	}
	if !validateNetCharUpdate(updatedNetCharList, "zone1-fog1-iperf", "ue1-iperf", 1, 1, 0, 500) {
		t.Fatalf("Error in Net Char update")
	}
	if !validateNetCharUpdate(updatedNetCharList, "zone1-fog1-svc", "ue2-svc", 1, 1, 0, 500) {
		t.Fatalf("Error in Net Char update")
	}

	// Verify algo calculations after Network Characteristic update
	fmt.Println("Update Net char")
	var netCharUpdateEvent dataModel.EventNetworkCharacteristicsUpdate
	netCharUpdateEvent.ElementName = "zone1-poa1"
	netCharUpdateEvent.ElementType = "POA"
	var netChar dataModel.NetworkCharacteristics
	netChar.Latency = 1          // no change
	netChar.LatencyVariation = 1 // no change
	netChar.PacketLoss = 0       // no change
	netChar.ThroughputDl = 100
	netCharUpdateEvent.NetChar = &netChar
	err = activeModel.UpdateNetChar(&netCharUpdateEvent)
	if err != nil {
		t.Fatalf("Error updating net char")
	}
	err = algo.ProcessScenario(activeModel)
	if err != nil {
		t.Fatalf("Failed to process scenario model")
	}
	updatedNetCharList = algo.CalculateNetChar()
	if len(updatedNetCharList) != 2 {
		t.Fatalf("Invalid net char update list")
	}
	if !validateNetCharUpdate(updatedNetCharList, "zone1-fog1-iperf", "ue1-iperf", 1, 1, 0, 50) {
		t.Fatalf("Error in Net Char update")
	}
	if !validateNetCharUpdate(updatedNetCharList, "zone1-fog1-svc", "ue2-svc", 1, 1, 0, 50) {
		t.Fatalf("Error in Net Char update")
	}

	// Update metrics & recalculate
	if !setMetrics(rc, "zone1-fog1-iperf", "ue1-iperf", 50) {
		t.Fatalf("Error updating metrics")
	}
	if !setMetrics(rc, "zone1-fog1-svc", "ue2-svc", 50) {
		t.Fatalf("Error updating metrics")
	}
	updatedNetCharList = algo.CalculateNetChar()
	if len(updatedNetCharList) != 0 {
		t.Fatalf("Invalid net char update list")
	}

	// Verify algo calculations after UE Mobility event update
	fmt.Println("Move ue1 to zone2-poa1")
	_, _, err = activeModel.MoveNode("ue1", "zone2-poa1")
	if err != nil {
		t.Fatalf("Error moving UE")
	}
	err = algo.ProcessScenario(activeModel)
	if err != nil {
		t.Fatalf("Failed to process scenario model")
	}

	// Update metrics & recalculate
	if !setMetrics(rc, "zone1-fog1-iperf", "ue1-iperf", 25) {
		t.Fatalf("Error updating metrics")
	}
	if !setMetrics(rc, "zone2-edge1-iperf", "ue1-iperf", 25) {
		t.Fatalf("Error updating metrics")
	}
	if !setMetrics(rc, "zone1-fog1-svc", "ue2-svc", 50) {
		t.Fatalf("Error updating metrics")
	}
	updatedNetCharList = algo.CalculateNetChar()
	if len(updatedNetCharList) != 19 {
		t.Fatalf("Invalid net char update list")
	}
	if !validateNetCharUpdate(updatedNetCharList, "zone1-fog1-iperf", "ue1-iperf", 41, 9, 0, 10) {
		t.Fatalf("Error in Net Char update")
	}
	if !validateNetCharUpdate(updatedNetCharList, "zone2-edge1-iperf", "ue1-iperf", 11, 3, 0, 10) {
		t.Fatalf("Error in Net Char update")
	}
	if !validateNetCharUpdate(updatedNetCharList, "zone1-fog1-svc", "ue2-svc", 1, 1, 0, 100) {
		t.Fatalf("Error in Net Char update")
	}

	// Update metrics & recalculate
	if !setMetrics(rc, "zone1-fog1-iperf", "ue1-iperf", 0) {
		t.Fatalf("Error updating metrics")
	}
	if !setMetrics(rc, "zone2-edge1-iperf", "ue1-iperf", 10) {
		t.Fatalf("Error updating metrics")
	}
	if !setMetrics(rc, "zone1-fog1-svc", "ue2-svc", 100) {
		t.Fatalf("Error updating metrics")
	}
	updatedNetCharList = algo.CalculateNetChar()
	if len(updatedNetCharList) != 2 {
		t.Fatalf("Invalid net char update list")
	}
	if !validateNetCharUpdate(updatedNetCharList, "zone1-fog1-iperf", "ue1-iperf", 41, 9, 0, 6) {
		t.Fatalf("Error in Net Char update")
	}
	if !validateNetCharUpdate(updatedNetCharList, "zone2-edge1-iperf", "ue1-iperf", 11, 3, 0, 20) {
		t.Fatalf("Error in Net Char update")
	}

	// Make sure we get no more updates when steady state is reached
	updatedNetCharList = algo.CalculateNetChar()
	if len(updatedNetCharList) != 0 {
		t.Fatalf("Invalid net char update list")
	}

	// Verify algo calculations after UE Mobility event update
	fmt.Println("Move ue1 to zone1-poa1")
	_, _, err = activeModel.MoveNode("ue1", "zone1-poa1")
	if err != nil {
		t.Fatalf("Error moving UE")
	}
	err = algo.ProcessScenario(activeModel)
	if err != nil {
		t.Fatalf("Failed to process scenario model")
	}

	// Update metrics & recalculate
	if !setMetrics(rc, "zone1-fog1-iperf", "ue1-iperf", 0) {
		t.Fatalf("Error updating metrics")
	}
	if !setMetrics(rc, "zone2-edge1-iperf", "ue1-iperf", 20) {
		t.Fatalf("Error updating metrics")
	}
	updatedNetCharList = algo.CalculateNetChar()
	if len(updatedNetCharList) != 19 {
		t.Fatalf("Invalid net char update list")
	}
	if !validateNetCharUpdate(updatedNetCharList, "zone2-edge1-iperf", "ue1-iperf", 41, 9, 0, 23) {
		t.Fatalf("Error in Net Char update")
	}
	if !validateNetCharUpdate(updatedNetCharList, "zone1-fog1-svc", "ue2-svc", 1, 1, 0, 77) {
		t.Fatalf("Error in Net Char update")
	}

	// Update metrics & recalculate
	if !setMetrics(rc, "zone1-fog1-iperf", "ue1-iperf", 23) {
		t.Fatalf("Error updating metrics")
	}
	if !setMetrics(rc, "zone2-edge1-iperf", "ue1-iperf", 0) {
		t.Fatalf("Error updating metrics")
	}
	if !setMetrics(rc, "zone1-fog1-svc", "ue2-svc", 77) {
		t.Fatalf("Error updating metrics")
	}
	updatedNetCharList = algo.CalculateNetChar()
	if len(updatedNetCharList) != 3 {
		t.Fatalf("Invalid net char update list")
	}
	if !validateNetCharUpdate(updatedNetCharList, "zone1-fog1-iperf", "ue1-iperf", 1, 1, 0, 26) {
		t.Fatalf("Error in Net Char update")
	}
	if !validateNetCharUpdate(updatedNetCharList, "zone2-edge1-iperf", "ue1-iperf", 41, 9, 0, 20) {
		t.Fatalf("Error in Net Char update")
	}
	if !validateNetCharUpdate(updatedNetCharList, "zone1-fog1-svc", "ue2-svc", 1, 1, 0, 74) {
		t.Fatalf("Error in Net Char update")
	}
}

func setMetrics(rc *redis.Connector, src string, dst string, throughput float64) bool {
	key := dkm.GetKeyRoot(testModuleNamespace) + metricsKey + dst + ":throughput"
	throughputStats := make(map[string]interface{})
	throughputStats[src] = throughput
	err := rc.SetEntry(key, throughputStats)
	return err == nil
}

func validatePath(algo *SegmentAlgorithm, flowName string, segmentCount int) bool {
	if flow, ok := algo.FlowMap[flowName]; ok {
		if flow.Path != nil || len(flow.Path.Segments) == segmentCount {
			return true
		}
	}
	return false
}

func validateNetCharUpdate(updatedNetCharList []FlowNetChar, src string, dst string, latency float64, jitter float64, packetloss float64, throughput float64) bool {
	found := false
	for _, flowNetChar := range updatedNetCharList {
		if flowNetChar.DstElemName == dst &&
			flowNetChar.SrcElemName == src &&
			flowNetChar.MyNetChar.Latency == latency &&
			flowNetChar.MyNetChar.Jitter == jitter &&
			flowNetChar.MyNetChar.PacketLoss == packetloss &&
			flowNetChar.MyNetChar.Throughput == throughput {

			found = true
			break
		}
	}
	return found
}
