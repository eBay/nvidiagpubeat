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
	"os/exec"
	"strconv"
	"strings"
)

//GPUCount provides interface to get gpu count command and run it.
type GPUCount interface {
	command() *exec.Cmd
	run(cmd *exec.Cmd, env string) (int, error)
}

//Count implements one flavour of GPUCount interface.
type Count struct {
}

//NewCount returns instance of Count
func newCount() Count {
	return Count{}
}

func (g Count) command() *exec.Cmd {
	cmd := "ls /dev | grep nvidia | grep -v nvidia-uvm | grep -v nvidiactl | grep -v modese | wc -l"
	return exec.Command("bash", "-c", cmd)
}

func (g Count) run(cmd *exec.Cmd, env string) (int, error) {
	if env == "test" {
		return 4, nil
	}
	out, err := cmd.Output()
	ret := 0
	if err == nil {
		ret, _ = strconv.Atoi(strings.TrimSpace(string(out)))
		return ret, nil
	} else {
		return -1, err
	}
}
