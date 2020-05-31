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
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
)

//GPUMetrics provides slice of metrics passed as argument for a given environment
type GPUMetrics interface {
	GetMetrics(env string, query string) ([]common.MapStr, error)
}

//Metrics implements one flavour of GPUMetrics interface.
type Metrics struct {
}

//NewMetrics returns instance of Metrics
func NewMetrics() Metrics {
	return Metrics{}
}

//Get return a slice of GPU metrics
func (m Metrics) Get(env string, query string) ([]common.MapStr, error) {
	gpuCount := newCount()
	gpuCountCmd := gpuCount.command()
	logp.Debug("nvidiagpubeat", "Determine number of gpu cards.")
	count, err := gpuCount.run(gpuCountCmd, env)
	if err != nil {
		logp.Err("Error %s . ", err)
		return nil, err
	}
	logp.Debug("nvidiagpubeat", "Number of gpu cards %d.", count)
	gpuUtilization := newUtilization()
	gpuUtilizationCmd := gpuUtilization.command(env, query)
	events, err := gpuUtilization.run(gpuUtilizationCmd, count, query, NewLocal())
	return events, err
}
