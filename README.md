<!--
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
-->

# nvidiagpubeat

Welcome to nvidiagpubeat.
nvidiagpubeat is an elastic beat that uses NVIDIA System Management Interface (nvidia-smi - https://developer.nvidia.com/nvidia-system-management-interface) to monitor NVIDIA GPU devices and can ingest metrics into Elastic search cluster. nvidia-smi is a command line utility, based on top of the NVIDIA Management Library (NVML), intended to aid in the management and monitoring of NVIDIA GPU devices.

nvidiagpubeat is built using Beats framework described at https://www.elastic.co/guide/en/beats/devguide/current/new-beat.html.

nvidiagpubeat elastic beat with help of nvidia-smi allows administrators to query GPU device state.  It is targeted at the TeslaTM, GRIDTM, QuadroTM and Titan X product, though limited support is also available on other NVIDIA GPUs.

NVIDIA-smi ships with NVIDIA GPU display drivers on Linux, and with 64bit Windows Server 2008 R2 and Windows 7.

nvidiagpubeat provides ability (look at nvidiagpubeat.yml) to configure metrics that needs to be monitored and by default it queries utilization.gpu,utilization.memory,memory.total,memory.free,memory.used,temperature.gpu,pstate and can ingest them into elastic search cluster for possibly visualization using Kibana.

## Getting Started with nvidiagpubeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7+
* virtualenv - https://virtualenv.pypa.io/en/stable/
```
sudo curl https://bootstrap.pypa.io/get-pip.py | sudo python
sudo pip install virtualenv
```
* glide - https://github.com/Masterminds/glide#install
```
brew install glide
```

### Init Project
To get running with nvidiagpubeat and also install the dependencies, run the following command:

```
#Use an empty directory
mkdir beats_dev


#Build with elastic beats branch=6.5 (Master branch might not always uptodate).
export WORKSPACE=`pwd`/beats_dev
export GOPATH=$WORKSPACE
git clone https://github.com/elastic/beats ${GOPATH}/src/github.com/elastic/beats --branch 6.5

#Clone nvidiagpubeat
mkdir $WORKSPACE/src/github.com/eBay
cd $WORKSPACE/src/github.com/eBay
git clone git@github.com:eBay/nvidiagpubeat.git

#Build
cd $WORKSPACE/src/github.com/ebay/nvidiagpubeat/
make setup
make
```
Above instructions will generate a binary in the same directory with the name nvidiagpubeat.

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes. To push nvidiagpubeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/eBay/nvidiagpubeat.git
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

Make changes (if any) and build the binary for nvidiagpubeat run the command below. This will generate a binary
in the same directory with the name nvidiagpubeat.

```
make
```

### Run

To run nvidiagpubeat in test environment with debugging output enabled, run:

```
export PATH=$PATH:.
./nvidiagpubeat -c nvidiagpubeat.yml -e -d "*"
```
nvidiagpubeat.yml contains env set to "test". This will run localnvidiasmi utility instead of nvidia-smi. Switch
env=production to use pre-installed nvidia-smi utility on a GPU machine.

Below is a sample event emitted by nvidiagpubeat.

```
Publish event: {
  "@timestamp": "2018-12-12T02:01:40.670Z",
  "@metadata": {
    "beat": "nvidiagpubeat",
    "type": "_doc",
    "version": "7.0.0"
  },
  "utilization": {
    "gpu": 4,
    "memory": 40
  },
  "temperature": {
    "gpu": 27
  },
  "host": {
    "name": "LM-SJL-11004865"
  },
  "pstate": 8,
  "@timestamp": "2018-12-12T02:01:40.669Z",
  "gpuIndex": 3,
  "type": "nvidiagpubeat",
  "memory": {
    "total": 6082,
    "free": 6082,
    "used": 0
  },
  "agent": {
    "type": "nvidiagpubeat",
    "hostname": "LM-SJL-11004865",
    "version": "7.0.0"
  }
}
```

### Test

To test nvidiagpubeat, run the following command:

```
make unit-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Cleanup dependencies
To remove them dependent libraries, builds, temporary files, files created during tests and the executatble, simply run

```
 rm -rf build/
 rm -rf data/
 rm -rf nvidiagpubeat
 rm -rf nvidiagpubeat.test
 rm -rf vendor
```

### Clone

To clone nvidiagpubeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/eBay
cd ${GOPATH}/github.com/eBay
git clone git@github.com:eBay/nvidiagpubeat.git
```
For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### License
Copyright 2016-2018 eBay Inc.
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
