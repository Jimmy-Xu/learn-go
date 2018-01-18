
# build

```
$ go build main.go
```

# usage

## available action

```
create-pod
get-pod
list-pod
update-pod
delete-pod

list-node

create-pv
get-pv
list-pv
delete-pv

create-pvc
get-pvc
list-pvc
delete-pvc
```

## read config from ~/.kube/config

```
//list pod
./main
or
./main --action list-pod

//create pod (default pod name is test-nginx)
./main --action create-pod

//list pod
./main --action list-pod

//get pod
./main --action get-pod

//update pod(change image)
./main --action update-pod

//delete pod
./main --action delete-pod

//list node
./main --action list-node

//specify pod name
./main --action create-pod --pod-name other-nginx

//specify apiserver
./main --action list-pod --server=http://127.0.0.1:8001

//show more log
./main --action list-pod -logtostderr -v=4
```

## use kube proxy
```
//start kube proxy
$ screen -S kubectl-proxy -L -d -m bash -c "kubectl proxy"

//run
$ ./main --action=list-pod --server=http://127.0.0.1:8001

//run with more log
$ ./main --action=list-pod --server=http://127.0.0.1:8001 -logtostderr -v=4
$ ./main --action=create-pod --server=http://127.0.0.1:8001 -logtostderr -v=4
```

## connect apirouter for gcp
```
export HYPER_ACCESS_KEY="0PN5J17HBGZHT7JJ3X82"
export HYPER_SECRET_KEY="uV3F3YluFJax1cknvbcGwgjvx4QpvB+leU8dUj2o"
export HYPER_REGION="gcp"

./main --action=list-pod --server=tcp://127.0.0.1:6443 -logtostderr -v=3
./main --action=create-pod --server=tcp://127.0.0.1:6443 -logtostderr -v=3
```

