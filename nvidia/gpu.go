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
	"encoding/csv"
	"errors"
	"io"
	"os/exec"
	"strconv"
	"strings"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

//GPUUtilization provides interface to utilization metrics and state of GPU.
type GPUUtilization interface {
	command(env string) *exec.Cmd
	run(cmd *exec.Cmd, gpuCount int, query string, action Action) ([]common.MapStr, error)
}

//Utilization implements one flavour of GPUCount interface.
type Utilization struct {
}

//newUtilization returns instance of Utilization
func newUtilization() Utilization {
	return Utilization{}
}

func (g Utilization) command(env string, query string) *exec.Cmd {
	if env == "test" {
		if strings.Contains(query, "query-compute-apps") {
			args := "query-compute-apps"
			return exec.Command("localnvidiasmi", args)
		} else {
			return exec.Command("localnvidiasmi")
		}
	}

	//For backward compatibility support the default query of --query-apps, if none is provided
	if strings.Contains(query, "=") {
		return exec.Command("nvidia-smi", query, "--format=csv")
	} else {
		return exec.Command("nvidia-smi", "--query-gpu="+query, "--format=csv")
	}
}

//Run the nvidiasmi command to collect GPU metrics
//Parse output and return events.
func (g Utilization) run(cmd *exec.Cmd, gpuCount int, query string, action Action) ([]common.MapStr, error) {
	logp.Info("Running command %s for query:  %s  with gpuCount %d", cmd, query, gpuCount)
	reader := action.start(cmd)
	gpuIndex := 0
	events := make([]common.MapStr, gpuCount, 2*gpuCount)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}
		// Ignore header
		if strings.Contains(line, "utilization") || strings.Contains(line, "gpu_name") || strings.Contains(line, "gpu_uuid") {
			continue
		}
		if len(line) == 0 {
			return nil, errors.New("Unable to fetch any events from nvidia-smi: Error " + err.Error())
		}

		// Remove units put by nvidia-smi
		line = strings.Replace(line, " %", "", -1)
		line = strings.Replace(line, " MiB", "", -1)
		line = strings.Replace(line, " MHz", "", -1)
		line = strings.Replace(line, " P", "", -1)
		line = strings.Replace(line, " W", "", -1)
		line = strings.Replace(line, " ", "", -1)

		r := csv.NewReader(strings.NewReader(line))
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		var headers []string
		if strings.Contains(query, "=") {
			rawHeaders := strings.Split(query, "=")
			headers = strings.Split(rawHeaders[1], ",")
		} else {
			headers = strings.Split(query, ",")
		}

		event := common.MapStr{
			"gpuIndex": gpuIndex,
			"type":     "nvidiagpubeat",
		}
		for i := 0; i < len(record) && i < len(headers); i++ {
			rawValue := record[i]
			iValue, err := strconv.ParseFloat(record[i], 64)
			if err == nil {
				event.Put(headers[i], iValue)
			} else {
				event.Put(headers[i], rawValue)
			}
		}
		events[gpuIndex] = event
		gpuIndex++
	}
	cmd.Wait()
	return events, nil
}
