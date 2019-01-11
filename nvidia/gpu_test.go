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

package nvidia

import (
	"os"
	"testing"
)

func Test_Command_TestEnv(t *testing.T) {
	util := newUtilization()
	cmd := util.command("test", "myquery")

	if len(cmd.Args) != 1 {
		t.Errorf("Expected %d, Actual %d", 1, len(cmd.Args))
	}

	if cmd.Args[0] != "localnvidiasmi" {
		t.Errorf("Expected %s, Actual %s", "localnvidiasmi", cmd.Args[0])
	}
}

func Test_Command_ProdEnv(t *testing.T) {
	util := newUtilization()
	cmd := util.command("prod", "myquery")

	if len(cmd.Args) != 3 {
		t.Errorf("Expected %d, Actual %d", 3, len(cmd.Args))
	}

	if cmd.Args[0] != "nvidia-smi" {
		t.Errorf("Expected %s, Actual %s", "nvidia-smi", cmd.Args[0])
	}

	if cmd.Args[1] != "--query-gpu=myquery" {
		t.Errorf("Expected %s, Actual %s", "--query-gpu=myquery", cmd.Args[1])
	}

	if cmd.Args[2] != "--format=csv" {
		t.Errorf("Expected %s, Actual %s", "--format=csv", cmd.Args[2])
	}
}

func Test_Run_TestEnv(t *testing.T) {
	util := newUtilization()
	query := "utilization.gpu,utilization.memory,memory.total,memory.free,memory.used,temperature.gpu,pstate"
	cmd := util.command("test", query)
	os.Setenv("PATH", ".")
	output, _ := util.run(cmd, 4, query, NewLocal())
	if output == nil {
		//TODO fix unit test case
		//t.Errorf("output cannot be nil")
	}
}

func Test_Run_ProdEnv(t *testing.T) {
	util := newUtilization()
	query := "utilization.gpu,utilization.memory,memory.total,memory.free,memory.used,temperature.gpu,pstate"
	cmd := util.command("prod", query)
	t.Logf("Command: %s", cmd.Path)
	output, _ := util.run(cmd, 4, query, MockLocal{})

	for _, o := range output {
		if o == nil {
			t.Errorf("output cannot be nil.")
		}
	}
}

func Test_Event_Contains_Type_Field(t *testing.T) {
	util := newUtilization()
	query := "utilization.gpu,utilization.memory,memory.total,memory.free,memory.used,temperature.gpu,pstate"
	cmd := util.command("prod", query)
	t.Logf("Command: %s", cmd.Path)

	output, _ := util.run(cmd, 4, query, MockLocal{})
	t.Logf("Nr. of Events: %d", len(output))

	for _, o := range output {
		if o["type"] != "nvidiagpubeat" {
			t.Errorf("event does not contain 'type' field equal to 'nvidiagpubeat'")
		}
	}
}
