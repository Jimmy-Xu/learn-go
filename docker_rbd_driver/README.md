## get go-ceph library

```
//install dependency
$ sudo apt-get install -y librados-dev librbd-dev


//download go-ceph
$ go get github.com/noahdesu/go-ceph/

//build docker source with rbd
$ ./docker.sh build v1.9.1
```


## Adjust rbd driver for docker 1.9.1

#### `docker/daemon/graphdriver/rbd/driver.go`

```
 import (
        log "github.com/Sirupsen/logrus"
        "github.com/docker/docker/daemon/graphdriver"
        "github.com/docker/docker/pkg/mount"
+       "github.com/docker/docker/pkg/idtools"
        "io/ioutil"

 type Driver struct {
        home string
        *RbdSet
+       uidMaps []idtools.IDMap
+       gidMaps []idtools.IDMap
 }

-func Init(home string, options []string) (graphdriver.Driver, error) {
+func Init(home string, options []string, uidMaps, gidMaps []idtools.IDMap) (graphdriver.Driver, error) {

-       rbdSet, err := NewRbdSet(home, true, options)
+       rbdSet, err := NewRbdSet(home, true, options, uidMaps, gidMaps)


```


#### `docker/daemon/graphdriver/rbd/rbd.go`
```
 import (
        "github.com/opencontainers/runc/libcontainer/label"
        "github.com/noahdesu/go-ceph/rados"
        "github.com/noahdesu/go-ceph/rbd"
+       "github.com/docker/docker/pkg/idtools"
        "os/exec"

 type RbdSet struct {
        filesystem   string
        mountOptions string
        mkfsArgs     []string
+       uidMaps      []idtools.IDMap
+       gidMaps      []idtools.IDMap
 }

-func NewRbdSet(root string, doInit bool, options []string) (*RbdSet, error) {
+func NewRbdSet(root string, doInit bool, options []string, uidMaps, gidMaps []idtools.IDMap) (*RbdSet, error) {

 func NewRbdSet(root string, doInit bool, options []string) (*RbdSet, error) {
                clientId:      "admin",
                configFile:    DefaultRadosConfigFile,
                filesystem:    "ext4",
+               uidMaps:       uidMaps,
+               gidMaps:       gidMaps,
 }
```

#### docker/`Makefile`
```
 binary: build
-       $(DOCKER_RUN_DOCKER) hack/make.sh binary
+       $(DOCKER_RUN_DOCKER) hack/make.sh dynbinary

```

#### `docker/Dockerfile`
```
+RUN apt-get install -y librbd-dev librados-dev
 COPY . /go/src/github.com/docker/docker

```