
# build

```
$ go build main.go
```

# usage

## read config from ~/.kube/config

```
//list pod
./main
or
./main --action list

//create pod
./main --action create
```

## use kube proxy
```
//start kube proxy
$ screen -S kubectl-proxy -L -d -m bash -c "kubectl proxy"

//run
$ ./main --action=list --host=http://127.0.0.1:8001

//run with more log
$ ./main --action=list --host=http://127.0.0.1:8001 -v=4 -logtostderr
$ ./main --action=create --host=http://127.0.0.1:8001 -v=4 -logtostderr
```

## connect apirouter for gcp
```
export HYPER_ACCESS_KEY="0PN5J17HBGZHT7JJ3X82"
export HYPER_SECRET_KEY="uV3F3YluFJax1cknvbcGwgjvx4QpvB+leU8dUj2o"
export HYPER_REGION="gcp"

./main --action=list --host=tcp://127.0.0.1:6443 -v=3 -logtostderr
./main --action=create --host=tcp://127.0.0.1:6443 -v=3 -logtostderr
```

