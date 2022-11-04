/*
 * Copyright (c) 2022  The AdvantEDGE Authors
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

package watchdog

import (
	"errors"
	"time"

	log "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger"
)

// Watchdog - Implements a Redis Watchdog
type Watchdog struct {
	name          string
	namespace     string
	peerName      string
	peerNamespace string
	isAlive       bool
	isStarted     bool
	pinger        *Pinger
	pingRate      time.Duration
	pongTimeout   time.Duration
	pongTime      time.Time
	ticker        *time.Ticker
}

// NewWatchdog - Create, Initialize and connect  a watchdog
func NewWatchdog(name string, namespace string, peerName string, peerNamespace string, dbAddr string) (w *Watchdog, err error) {
	if name == "" {
		err = errors.New("Missing watchdog name")
		log.Error(err)
		return nil, err
	}

	w = new(Watchdog)
	w.name = name
	w.namespace = namespace
	w.peerName = peerName
	w.peerNamespace = peerNamespace
	w.isStarted = false

	w.pinger, err = NewPinger(w.name, namespace, dbAddr)
	if err != nil {
		log.Error("Error creating watchdog: ", err)
		return nil, err
	}
	log.Debug("Watchdog created ", w.name)
	return w, nil
}

// Start - starts watchdog monitoring
func (w *Watchdog) Start(rate time.Duration, timeout time.Duration) (err error) {
	w.isStarted = true
	w.isAlive = true // start with a positive attitude!
	w.pongTime = time.Now()
	w.pingRate = rate
	w.pongTimeout = timeout
	w.ticker = time.NewTicker(w.pingRate)

	err = w.pinger.Start()
	if err != nil {
		log.Error("Watchdog failed to start pinger ", w.name)
		log.Error(err)
		return err
	}

	go w.watchdogTask()
	log.Debug("Watchdog started ", w.name, "(rate=", w.pingRate, ", timeout=", w.pongTimeout, ")")
	return nil
}

func (w *Watchdog) watchdogTask() {
	log.Debug("Watchdog task started: ", w.name)
	for range w.ticker.C {
		isAlive := w.pinger.Ping(w.peerName, w.peerNamespace, time.Now().String())
		if isAlive {
			w.pongTime = time.Now()
		}
	}
	log.Debug("Watchdog task terminated: ", w.name)
}

// Stop - stops watchdog monitoring
func (w *Watchdog) Stop() {
	if w.isStarted {
		w.ticker.Stop()
		_ = w.pinger.Stop()
		log.Debug("Watchdog stopped ", w.name)
	}
}

// IsAlive - Indicates if the monitored resource is alive
func (w *Watchdog) IsAlive() bool {
	if time.Since(w.pongTime) > w.pongTimeout {
		w.isAlive = false
	} else {
		w.isAlive = true
	}
	return w.isAlive
}
