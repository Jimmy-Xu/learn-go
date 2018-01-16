
# build

```
$ go build main.go
```

# usage

## read config from ~/.kube/config

```
//run
./main
```

## use kube proxy
```
//start kube proxy
$ screen -S kubectl-proxy -L -d -m bash -c "kubectl proxy"

//run
$ ./main --host=http://127.0.0.1:8001

```
