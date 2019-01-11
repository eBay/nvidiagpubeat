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

import "testing"

func Test_GPUCount_Command(t *testing.T) {
	count := newCount()
	cmd := count.command()

	if len(cmd.Args) != 3 {
		t.Errorf("Expected %d, Actual %d", 3, len(cmd.Args))
	}

	if cmd.Args[0] != "bash" {
		t.Errorf("Expected %s, Actual %s", "bash", cmd.Args[0])
	}

	if cmd.Args[1] != "-c" {
		t.Errorf("Expected %s, Actual %s", "bash", cmd.Args[1])
	}

	if cmd.Args[2] != "ls /dev | grep nvidia | grep -v nvidia-uvm | grep -v nvidiactl | wc -l" {
		t.Errorf("Expected %s, Actual %s", "bash", cmd.Args[2])
	}
}

func Test_GPUCount_Run_TestEnv(t *testing.T) {
	count := newCount()
	cmd := count.command()
	gpuCount, _ := count.run(cmd, "test")

	if gpuCount != 4 {
		t.Errorf("Expected %d, Actual %d", 4, gpuCount)
	}
}

func Test_GPUCount_Run_ProdEnv(t *testing.T) {
	count := newCount()
	cmd := count.command()
	gpuCount, _ := count.run(cmd, "prod")

	if gpuCount != -1 {
		t.Errorf("Expected %d, Actual %d", -1, gpuCount)
	}
}
