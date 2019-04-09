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

	"github.com/armPelionEdge/maestro-plugins-template/teststructs"
	"github.com/armPelionEdge/maestroSpecs"
)

var mylogger maestroSpecs.Logger

// InitMaestroPlugin is always called on plugin startup one time only.
func InitMaestroPlugin(opts *maestroSpecs.PluginOpts, api *maestroSpecs.API, log maestroSpecs.Logger) (err error) {
	mylogger = log
	//	fmt.Printf("InitMaestroPlugin() dummy\n")
	// _watchdog = new(watchdog)
	// Watchdog = _watchdog
	return
}

// CallMePeriodically is a periodically called test function
// Must meet this call spec: func(num int) (contin bool, err error)
func CallMePeriodically(num int, param *maestroSpecs.CommonParam) (contin bool, err error) {
	fmt.Printf("In CallMePeriodically %d\n", num)
	mylogger.Debugf("In CallMePeriodically %d", num)
	if num < 10 {
		contin = true
	}
	if num == 5 {
		mylogger.Errorf("CallMePeriodically generating error %d", num)
		err = errors.New("Some error")
	}

	if param.Param == nil {
		param.Param = new(teststructs.Stuff)
	}

	stuff, ok := param.Param.(*teststructs.Stuff)
	if ok {
		stuff.PanicAtTheDisco = num
	}

	return
}
