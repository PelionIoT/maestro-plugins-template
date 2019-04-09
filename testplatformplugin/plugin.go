package main

// Copyright (c) 2018, Arm Limited and affiliates.
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"errors"
	"fmt"

	"github.com/WigWagCo/maestroSpecs"
	"github.com/WigWagCo/maestroSpecs/templates"
)

const (
	internalInterval = 15 // seconds
	stop             = 1
	keepalive        = 2
)

// This is a Go plugin, meeting the specs of maestroSpecs/watchdog.go

// type watchdog struct {
// 	ready    bool
// 	log      maestroSpecs.WatchdogLogger
// 	ctrlChan chan int
// }

// // Watchdog is the single instance of this watchdog
// var _watchdog *watchdog

// // Watchdog this is the exported interface of the watchdog
// var Watchdog maestroSpecs.Watchdog

//var Watchdog maestroSpecs.Watchdog

// InitMaestroPlugin is always called by Maestro any time it inits a plugin
func InitMaestroPlugin(opts *maestroSpecs.PluginOpts, api *maestroSpecs.API, log maestroSpecs.Logger) (err error) {
	//	fmt.Printf("InitMaestroPlugin() dummy\n")
	// _watchdog = new(watchdog)
	// Watchdog = _watchdog
	return
}

type eeprom struct {
	BRAND              string `dict:"BRAND"`
	DEVICE             string `dict:"DEVICE"`
	UUID               string `dict:"BRAND"`
	RELAYID            string `json:"relayID" dict:"RELAYID"`
	HARDWARE_VERSION   string `json:"hardwareVersion" dict:"HARDWARE_VERSION"`
	WW_PLATFORM        string `dict:"WW_PLATFORM"`
	FIRMWARE_VERSION   string `dict:"FIRMWARE_VERSION"`
	RADIO_CONFIG       string `json:"radioConfig" dict:"RADIO_CONFIG"`
	YEAR               string `json:"year" dict:"YEAR"`
	MONTH              string `json:"month" dict:"MONTH"`
	BATCH              string `json:"batch" dict:"BATCH"`
	ETHERNET_MAC       string `dict:"ETHERNET_MAC"`
	SIXLBR_MAC         string `dict:"SIXLBR_MAC"`
	ETHERNET_MAC_ARRAY []byte `json:"ethernetMAC" dict:"ETHERNET_MAC_ARRAY"`
	SIXLBR_MAC_ARRAY   []byte `json:"sixBMAC" dict:"SIXLBR_MAC_ARRAY"`
	RELAY_SECRET       string `json:"relaySecret" dict:"RELAY_SECRET"`
	PAIRING_CODE       string `json:"pairingCode" dict:"PAIRING_CODE"`
	LED_CONFIG         string `json:"ledConfig" dict:"LED_CONFIG"`
	LED_COLOR_PROFILE  string `json:"ledConfig" dict:"LED_COLOR_PROFILE"`
	CLOUD_URL          string `json:"cloudURL" dict:"CLOUD_URL"`
	CLOUD_DEVJS_URL    string `json:"devicejsCloudURL" dict:"CLOUD_DEVJS_URL"`
	CLOUD_DDB_URL      string `json:"devicedbCloudURL" dict:"CLOUD_DDB_URL"`
	CLOUD_DDB_URL_RES  string `dict:"CLOUD_DDB_URL_RES"`
}

func GetPlatformVars(dict *templates.TemplateVarDictionary, log maestroSpecs.Logger) (err error) {
	//Read eeprom
	// eepromData, err := read_eeprom()
	// if err != nil {
	// 	return
	// }

	if dict == nil {
		err = errors.New("no dictionary")
		return
	}

	//Write certs
	//	pathErr := os.Mkdir("/wigwag/devicejs-core-modules/Runner/.ssl", os.ModePerm)

	// err = ioutil.WriteFile("/wigwag/devicejs-core-modules/Runner/.ssl/ca.cert.pem", []byte(eepromData.SSL.Ca.Ca), 0644)
	// if err != nil {
	// 	return
	// }
	dict.AddArch("SSL_CA_PEM", "SSL-CA-DATA-HERE")
	// err = ioutil.WriteFile("/wigwag/devicejs-core-modules/Runner/.ssl/intermediate.cert.pem", []byte(eepromData.SSL.Ca.Intermediate), 0644)
	// if err != nil {
	// 	return
	// }
	dict.AddArch("SSL_CA_INTERMEDIATE_PEM", "SSL-CA-INTERMEDIATE-DATA-HERE")
	// err = ioutil.WriteFile("/wigwag/devicejs-core-modules/Runner/.ssl/client.key.pem", []byte(eepromData.SSL.Client.Key), 0644)
	// if err != nil {
	// 	return
	// }
	dict.AddArch("SSL_CLIENT_KEY_PEM", "SSL-CLIENT-KEY-DATA-HERE")
	// err = ioutil.WriteFile("/wigwag/devicejs-core-modules/Runner/.ssl/client.cert.pem", []byte(eepromData.SSL.Client.Certificate), 0644)
	// if err != nil {
	// 	return
	// }
	dict.AddArch("SSL_CLIENT_CERT_PEM", "SSL-CLIENT-CERT-DATA-HERE")
	// err = ioutil.WriteFile("/wigwag/devicejs-core-modules/Runner/.ssl/server.key.pem", []byte(eepromData.SSL.Server.Key), 0644)
	// if err != nil {
	// 	return
	// }
	dict.AddArch("SSL_SERVER_KEY_PEM", "SSL-SERVER-KEY-DATA-HERE")
	// err = ioutil.WriteFile("/wigwag/devicejs-core-modules/Runner/.ssl/server.cert.pem", []byte(eepromData.SSL.Server.Certificate), 0644)
	// if err != nil {
	// 	return
	// }
	dict.AddArch("SSL_SERVER_CERT_PEM", "SSL-SERVER-CERT-DATA-HERE")
	// err = ioutil.WriteFile("/wigwag/devicejs-core-modules/Runner/.ssl/ca-chain.cert.pem", []byte(eepromData.SSL.Ca.Ca+eepromData.SSL.Ca.Intermediate), 0644)
	// if err != nil {
	// 	return
	// }
	dict.AddArch("SSL_CA_CHAIN_PEM", "SSL-CA-CHAIN-DATA-HERE")

	// just test data
	eepromData := &eeprom{
		CLOUD_DDB_URL:     "https://devcloud-devicedb.wigwag.io",
		CLOUD_DDB_URL_RES: "https://devcloud-devicedb.wigwag.io",
		CLOUD_DEVJS_URL:   "https://devcloud-devicejs.wigwag.io",
		RELAYID:           "WWTEST001",
		WW_PLATFORM:       "test-platform",
		PAIRING_CODE:      "123",
		RELAY_SECRET:      "verysecret",
	}

	var n int
	// this adds in all the rest of the struct fields with tags 'dict'
	// you can pass a struct or a pointer to a struct
	n, err = dict.AddTaggedStructArch(eepromData)

	fmt.Printf("PLATFORM_RP200: %d values\n", n)

	if err != nil {
		log.Errorf("Failed to add struct value: %s", err.Error())
	}

	// to test errors being reported.
	// err = errors.New("OUCH!!")

	// put all these found vars into the dictionary
	// for _, eeprom_entry := range eepromData {
	// 	dict.AddArch(eeprom_entry.name, eeprom_entry.data)
	// }
	return
}

// GeneratePlatformDeviceKeyNCert generates key and cert for the platform
func GeneratePlatformDeviceKeyNCert(dict *templates.TemplateVarDictionary, deviceid string, accountid string, log maestroSpecs.Logger) (key string, cert string, err error) {
	key = "KEY:" + deviceid + ":" + accountid
	cert = "CERT:" + deviceid + ":" + accountid
	log.Debugf("GeneratePlatformDeviceKeyNCert() done\n")
	return
}

// WritePlatformDeviceKeyNCert writes a key and cert for the platform
func WritePlatformDeviceKeyNCert(dict *templates.TemplateVarDictionary, key string, cert string, log maestroSpecs.Logger) (err error) {
	fmt.Printf("WRITE KEY:%s\n", key)
	fmt.Printf("WRITE CERT:%s\n", cert)
	log.Debugf("WritePlatformDeviceKeyNCert() done")
	return
}

// func (wd *watchdog) watchdogRunner() {
// 	if wd.ready {
// 		wd.log.Errorf("Watchdog (dummy): ERROR - the watchdog was enabled twice.")
// 		return
// 	}
// 	wd.ready = true
// 	interval := internalInterval * time.Second
// wdLoop:
// 	for {
// 		wd.log.Debugf("Watchdog (dummy): top of wdLoop.")
// 		select {
// 		case code := <-wd.ctrlChan:
// 			switch code {
// 			case stop:
// 				break wdLoop
// 			default:
// 				interval = internalInterval * time.Second
// 				continue
// 			}
// 		case <-time.After(interval):
// 			wd.log.Errorf("Watchdog (dummy): WATCHDOG TIMED OUT!!! SHOULD NOT HAPPEN.")
// 			interval = 2 * time.Second
// 		}
// 	}
// 	wd.log.Debugf("Watchdog (dummy): stopped")
// 	wd.ready = false
// }

// // Called by Maestro upon load. If the watchdog needs a setup procedure this
// // should occur here. An error returned will prevent further calls from Maestro
// // including KeepAlive()
// func (wd *watchdog) Setup(config *maestroSpecs.WatchdogConfig, logger maestroSpecs.WatchdogLogger) (err error) {
// 	wd.log = logger
// 	wd.log.Debugf("Watchdog (dummy) Got Setup() call with config: %+v", config)
// 	wd.ctrlChan = make(chan int)
// 	return
// }

// // func SetupWatchdog(config *maestroSpecs.WatchdogConfig, logger maestroSpecs.WatchdogLogger) (err error, watchdog maestroSpecs.Watchdog) {
// // 	wd = new(watchdog)
// // 	wd.log = logger
// // 	wd.ready = true
// // 	wd.log.Debugf("Got Setup() call with config: %+v", config)
// // 	wd.ctrlChan = make(chan int)
// // 	return
// // }

// // CriticalInterval returns the time interval the watchdog *must* be called
// // in order to keep the system alive. The implementer should build in enough buffer
// // to this for the plugin to do its work of keep any hw / sw watchdog alive.
// // Maestro will call KeepAlive() at this interval or quicker, but never slower.
// func (wd *watchdog) CriticalInterval() time.Duration {
// 	// 5 seconds buffer please
// 	return time.Second * (internalInterval - 5)
// }

// // KeepAlive is called by Maestro to keep the watch dog up. A call of KeepAlive
// // means the watchdog plug can safely assume all is normal on the system
// func (wd *watchdog) KeepAlive() (err error) {
// 	select {
// 	case wd.ctrlChan <- keepalive:
// 		wd.log.Debugf("Watchdog (dummy): KeepAlive() called")
// 	default:
// 		wd.log.Errorf("Watchdog (dummy): KeepAlive() called - but would block. wd not running?")
// 	}
// 	return
// }

// // NotOk is called by Maestro at the same interval as KeepAlive, but in leiu of it
// // when the system is not meeting the criteria to keep the watch dog up
// func (wd *watchdog) NotOk() (err error) {
// 	wd.log.Errorf("Watchdog (dummy) Not OK.")
// 	return
// }

// // Disabled is called when Maestro desires, typically from a command it recieved, to
// // disabled the watchdog. Usually for debugging. Implementors should disable the watchdog
// // or if this is impossible, they can return an error.
// func (wd *watchdog) Disable() (err error) {
// 	if wd.ready {
// 		select {
// 		case wd.ctrlChan <- stop:
// 			wd.log.Debugf("Watchdog (dummy): Disable() called")
// 		default:
// 			wd.log.Errorf("Watchdog (dummy): Disable() called - but would block. wd not running?")
// 		}
// 	}
// 	return
// }

// // Enable() is not called normally. But if Disable() is called, the a call of Enable()
// // should renable the watchdog if it was disabled.
// func (wd *watchdog) Enable() (err error) {
// 	wd.log.Debugf("Watchdog (dummy): Enable() called")
// 	go wd.watchdogRunner()
// 	return
// }

// go build -buildmode=plugin
