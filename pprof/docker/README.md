# Generate callgraph for dokcer with pprof
## Add pprof to docker.go and build docker

```
$ ./docker.sh build
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
