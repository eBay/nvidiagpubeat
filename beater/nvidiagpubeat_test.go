//******************************************************************
//Copyright 2018 eBay Inc.
//Architect/Developer: Deepak Vasthimal

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at

// https://www.apache.org/licenses/LICENSE-2.0

//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//******************************************************************

package beater

import (
	"testing"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
)

func Test_New(t *testing.T) {
	bt := &beat.Beat{}
	cfg := common.NewConfig()

	nvBeat, err := New(bt, cfg)
	if nvBeat == nil {
		t.Errorf("nvBeat instantiation failed")
	}

	if err != nil {
		t.Errorf("nvBeat instantiation failed, Error : %s", err)
	}
}

func Test_Run(t *testing.T) {
	bt := &beat.Beat{}
	cfg := common.NewConfig()

	nvBeat, err := New(bt, cfg)
	if nvBeat == nil {
		t.Errorf("nvBeat instantiation failed")
	}

	if err != nil {
		t.Errorf("nvBeat instantiation failed, Error : %s", err)
	}

	//TODO Instantiate bt and test Run()
	//nvBeat.Run(bt)
}

func Test_Stop(t *testing.T) {
	bt := &beat.Beat{}
	cfg := common.NewConfig()

	nvBeat, err := New(bt, cfg)
	if nvBeat == nil {
		t.Errorf("nvBeat instantiation failed")
	}

	if err != nil {
		t.Errorf("nvBeat instantiation failed, Error : %s", err)
	}

	//TODO Instantiate bt and test Run()
	//nvBeat.Stop()
}
