
> REF: 
- **rbd driver for docker**
http://hustcat.github.io/run-docker-on-ceph/
https://github.com/hustcat/docker
https://github.com/hustcat/docker-graph-driver
- **create ceph cluster(vagrant+libvirt+kvm+centos7)**
https://github.com/Jimmy-Xu/my-vagrant/tree/master/ceph


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

## run docker daemon with rbd driver

```
// prepare rbd device
$ sudo ceph osd pool create test_pool 4096
$ sudo rbd create test_pool/test_image --size 4096
$ sudo rbd map test_pool/test_image
$ sudo rbd showmapped
$ sudo mkfs.ext4 /dev/rbd0
$ sudo mkdir -p /var/lib/docker-rbd
$ sudo mount /dev/rbd0 /var/lib/docker-rbd
$ df -hT

// start docker daemon
$ sudo service docker stop
$ sudo /home/xjimmy/gopath/src/github.com/docker/docker/bundles/latest/dynbinary/docker daemon -D -s rbd -g /var/lib/docker-rbd

$ docker info | grep rbd
$ docker pull busybox
$ docker run busybox uname -a
```

## FAQ

#### Q1 Missing `libdevmapper.so.1.02` when start docker daemon
``` 
// Error message
$ sudo /home/xjimmy/gopath/src/github.com/docker/docker/bundles/latest/dynbinary/docker daemon
/home/xjimmy/gopath/src/github.com/docker/docker/bundles/latest/dynbinary/docker: error while loading shared libraries: libdevmapper.so.1.02: cannot open shared object file: No such file or directory

// Solution
$ find /lib -name "libdevmapper.so*"
  /lib/x86_64-linux-gnu/libdevmapper.so.1.02.1
$ sudo ln -s /lib/x86_64-linux-gnu/libdevmapper.so.1.02.1 /lib/x86_64-linux-gnu/libdevmapper.so.1.02 && sudo ldconfig 
```