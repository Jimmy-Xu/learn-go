# Generate callgraph for hpyerd with pprof
## Add pprof to hyperd.go and build hyper

```
$ ./hyper.sh build
```

## Run hyperd with `-cpuprofile`

```
$ ./hyper.sh run
```

## Pull image with hyper cli

```
$ hyper pull busybox
```

## Stop hyperd
just press `ctrl + c`

## Generate callgraph
> convert `/tmp/hyperd_cpu.pprof` to `~/hyperd_callgraph.pdf`

```
$ ./hyper.sh graph
```

## View result
open `~/hyperd_callgraph.pdf`
