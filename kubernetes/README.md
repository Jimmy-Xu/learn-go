
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

//run with more log
$ ./main --host=http://127.0.0.1:8001 -v=4 -logtostderr
```

## connect apirouter for gcp
```
export HYPER_ACCESS_KEY="0PN5J17HBGZHT7JJ3X82"
export HYPER_SECRET_KEY="uV3F3YluFJax1cknvbcGwgjvx4QpvB+leU8dUj2o"
export HYPER_REGION="gcp"

./main --host=tcp://127.0.0.1:6443 -v=4 -logtostderr

```
