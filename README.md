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
mkdir $WORKSPACE/src/github.com/ebay
cd $WORKSPACE/src/github.com/ebay
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

### Run in production

To run nvidiagpubeat with pre-installed nvidia-smi that is available in PATH, switch to "test" environment 
in nvidiagpubeat.yml and use 

```yaml
env: "production"
```

```bash
export PATH=$PATH:.
./nvidiagpubeat -c nvidiagpubeat.yml -e -d "*" -E seccomp.enabled=false
```

`seccomp.enabled` setting : nvidiagpubeat uses libbeat framework. For security purposes the libbeat framework by default 
drops the ability to fork/exec. As nvidiagpubeat executes `nvidia-smi`, security setting must be disabled by 
setting through command line.

### Run in test (macOS)
To run nvidiagpubeat with pre-packaged `localnvidiasmi` switch to "test" environment in nvidiagpubeat.yml and use
```yaml
env: "test"
```

```bash
export PATH=$PATH:.
./nvidiagpubeat -c nvidiagpubeat.yml -e -d "*" -E seccomp.enabled=false
```

localnvidiasmi executable built for macOS and is a mock GPU event generator. 


### Sample event
Below is a sample event emitted by nvidiagpubeat.

```
Publish event: {
  "@timestamp": "2019-02-07T00:39:08.814Z",
  "@metadata": {
    "beat": "nvidiagpubeat",
    "type": "doc",
    "version": "6.5.5"
  },
  "gpuIndex": 3,
  "memory": {
    "total": 11448,
    "free": 479,
    "used": 10969
  },
  "host": {
    "name": "hostname.company.com"
  },
  "temperature": {
    "gpu": 64
  },
  "pstate": 0,
  "@timestamp": "2019-02-07T00:39:08.811Z",
  "type": "nvidiagpubeat",
  "utilization": {
    "gpu": 15,
    "memory": 12
  },
  "beat": {
    "version": "6.5.5",
    "name": "hostname.company.com",
    "hostname": "hostname.company.com"
  }
}
```

### Test cases

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
mkdir -p ${GOPATH}/github.com/ebay
cd ${GOPATH}/github.com/ebay
git clone git@github.com:eBay/nvidiagpubeat.git
```
For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


### Build nvidiagpubeat on RedHat EL7

Here below the speps to build nvidiagpubeat binary on a RedHat EL7 system


#### Software requirements installation

```
yum install python-virtualenv
yum install golang
```
#### Actual build

```
cd ~
mkdir beats_dev
export WORKSPACE=`pwd`/beats_dev
export GOPATH=$WORKSPACE
git clone https://github.com/elastic/beats ${GOPATH}/src/github.com/elastic/beats --branch 6.5
mkdir -p $WORKSPACE/src/github.com/ebay
cd $WORKSPACE/src/github.com/ebay/
git clone https://github.com/eBay/nvidiagpubeat.git
cd nvidiagpubeat
make copy-vendor
make
ls -l nvidiagpubeat
-rwxr-xr-x 1 user user 36862632  7 feb 09.04 nvidiagpubeat
```

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
