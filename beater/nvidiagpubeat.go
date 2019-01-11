/***********************************************************************
This module was automatically generated using the framework found below:
https://www.elastic.co/guide/en/beats/devguide/current/new-beat.html

Modifications to auto-generated code - Copyright 2018 eBay Inc.
Architect/Developer: Deepak Vasthimal

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
************************************************************************/

package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/ebay/nvidiagpubeat/config"
	"github.com/ebay/nvidiagpubeat/nvidia"
)

type Nvidiagpubeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

//New Creates the Beat object
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Nvidiagpubeat{
		done:   make(chan struct{}),
		config: config,
	}
	return bt, nil
}

//Run Contains the main application loop that captures data and sends it to the defined output using the publisher
func (bt *Nvidiagpubeat) Run(b *beat.Beat) error {
	logp.Info("nvidiagpubeat is running for ** %s ** environment. ! Hit CTRL-C to stop it.", bt.config.Env)

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}
	ticker := time.NewTicker(bt.config.Period)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		metrics := nvidia.NewMetrics()
		events, err := metrics.Get(bt.config.Env, bt.config.Query)
		if err != nil {
			logp.Err("Event not generated, error: %s", err.Error())
		} else {
			logp.Debug("nvidiagpubeat", "Event generated, Attempting to publish to configured output.")
			for _, gpuevent := range events {
				if gpuevent != nil {
					event := beat.Event{
						Timestamp: time.Now(),
						Fields:    gpuevent,
					}
					bt.client.Publish(event)
				}
			}
		}
	}
}

//Stop Contains logic that is called when the Beat is signaled to stop
func (bt *Nvidiagpubeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
