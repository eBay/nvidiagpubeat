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
	"bufio"
	"os/exec"

	"github.com/elastic/beats/libbeat/logp"
)

//Action provides interface to start execution of a command
type Action interface {
	start(cmd *exec.Cmd) *bufio.Reader
}

//Local implements one flavour of Action interface.
type Local struct {
}

//NewLocal returns instance of Local
func NewLocal() Local {
	return Local{}
}

//Start starts cmd and returns reader object that contains output of command
func (local Local) start(cmd *exec.Cmd) *bufio.Reader {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logp.Err("Error executing cmd, error: %s", err.Error())
	}
	cmd.Start()
	return bufio.NewReader(stdout)
}
