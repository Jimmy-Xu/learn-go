# Generate callgraph for dokcer with pprof
## Add pprof to docker.go and build docker

```
$ ./docker.sh build
or
$ ./docker.sh build v1.9.1
or
$ ./docker.sh build v1.10.0-rc1
or
$ ./docker.sh build master

// build result
$ docker images | grep docker-dev
  docker-dev:heads/v1.10.0-rc1      latest      505c47fa88c3      35 seconds ago     1.871 GB
  docker-dev:heads/v1.9.1           latest      7a444dd07ae9      23 minutes ago     1.935 GB
  docker-dev                        master      57a8ae00bb34      41 minutes ago     1.949 GB
```

## Run docker with `--cpuprofile`

```
$ ./docker.sh run
```

## Pull image with docker cli

```
$ docker pull busybox
```

## Stop docker
just press `ctrl + c`

## Generate callgraph
> convert `/tmp/docker_cpu.pprof` to `~/docker_callgraph.pdf`

```
$ ./docker.sh graph
```

## View result
open `~/docker_callgraph.pdf`
