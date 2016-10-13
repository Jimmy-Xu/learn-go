websocket client for hyper
===========================

Example client for connect to websocket api `/events/ws` of Hyper.sh apirouter

- SSL: `InsecureSkipVerify: true`
- Sign4: Use util/sign4.go

# Usage

## start apirouter
https://github.com/getdvm/hyper-api-router/pull/298
```
$ make
$ ./apirouter --insecure-bind-address=0.0.0.0 --trust-sources=0.0.0.0/0 --log_dir=/var/log/apirouter -v=4 --max-requests-inflight=0 --mongodb-url 147.x5.x5.x --volume-init-tenant=4a3300e7094f4c76a2cxxxxxxxxxxc6a
```

## run client


### prepare
```
$ go get github.com/Jimmy-Xu/learn-go/websocket/gorilla/practice/hyper
$ cd $GOPATH/src/github.com/Jimmy-Xu/learn-go/websocket/gorilla/practice/hyper
$ go run hyper-client.go --help   
Usage of /tmp/go-build064378304/command-line-arguments/_obj/exe/hyper-client:
  -accessKey string
    	hyper access key
  -addr string
    	apirouter entrypoint (default "147.x5.x5.x7:6443")
  -filter value
    	filter event by container,image,label,event (default string method)
  -pretty
    	pretty print result
  -secretKey string
    	hyper secret key
```

### Watch all events

`accessKey` and `secretKey` are required options

```
$ go run hyper-client.go --accessKey KXARxxxxx5WR8 --secretKey Ema5xxxxxxxxxZgNe
connecting to wss://147.x5.x5.x7:6443/events/ws
connected, watching event now:
{"status":"start","id":"f29698cac3f6f66e84790fb12b3e5e4f3455b89b3ff12150ac4d86b8b90d9179","from":"xjimmyshcn/busybox","Type":"container","Action":"start","Actor":{"ID":"f29698cac3f6f66e84790fb12b3e5e4f3455b89b3ff12150ac4d86b8b90d9179","Attributes":{"":"","exitCode":"0","image":"xjimmyshcn/busybox","name":"test4","sh_hyper_instancetype":"s4","test1":"","test2":"test2","test3":"test3=test3"}},"time":1476375774,"timeNano":1476375774255155116}
{"status":"stop","id":"f29698cac3f6f66e84790fb12b3e5e4f3455b89b3ff12150ac4d86b8b90d9179","from":"xjimmyshcn/busybox","Type":"container","Action":"stop","Actor":{"ID":"f29698cac3f6f66e84790fb12b3e5e4f3455b89b3ff12150ac4d86b8b90d9179","Attributes":{"":"","exitCode":"0","image":"xjimmyshcn/busybox","name":"test4","sh_hyper_instancetype":"s4","test1":"","test2":"test2","test3":"test3=test3"}},"time":1476375778,"timeNano":1476375778304732322}
```

### Output pretty json

use option `--pretty`

```
$ go run hyper-client.go --addr=147.x5.x5.x7:6443 --accessKey KXARxxxxx5WR8 --secretKey Ema5xxxxxxxxxZgNe --pretty
connecting to wss://147.x5.x5.x7:6443/events/ws
connected, watching event now:
{
  "Action": "start",
  "Actor": {
    "Attributes": {
      "": "",
      "exitCode": "0",
      "image": "xjimmyshcn/busybox",
      "name": "test4",
      "sh_hyper_instancetype": "s4",
      "test1": "",
      "test2": "test2",
      "test3": "test3=test3"
    },
    "ID": "f29698cac3f6f66e84790fb12b3e5e4f3455b89b3ff12150ac4d86b8b90d9179"
  },
  "Type": "container",
  "from": "xjimmyshcn/busybox",
  "id": "f29698cac3f6f66e84790fb12b3e5e4f3455b89b3ff12150ac4d86b8b90d9179",
  "status": "start",
  "time": 1.476375852e+09,
  "timeNano": 1.4763758521916593e+18
}
```

### Watch event with filter

use option `--filter`, support filter by `container,image,label,event`
- **container**: container id or name
- **image**: imageid or name
- **label**: label of container
- **event**: `start|stop`

```
$ go run hyper-client.go --addr=147.x5.x5.x7:6443 --accessKey KXARxxxxx5WR8 --secretKey Ema5xxxxxxxxxZgNe --filter=container=test4,image=xjimmyshcn/busybox,event=stop,label=test1,label=test2=test2
connecting to wss://147.x5.x5.x7:6443/events/ws?filters={"container":{"test4":true},"event":{"stop":true},"image":{"xjimmyshcn/busybox":true},"label":{"test1":true,"test2=test2":true}}
connected, watching event now:
{"status":"stop","id":"f29698cac3f6f66e84790fb12b3e5e4f3455b89b3ff12150ac4d86b8b90d9179","from":"xjimmyshcn/busybox","Type":"container","Action":"stop","Actor":{"ID":"f29698cac3f6f66e84790fb12b3e5e4f3455b89b3ff12150ac4d86b8b90d9179","Attributes":{"":"","exitCode":"0","image":"xjimmyshcn/busybox","name":"test4","sh_hyper_instancetype":"s4","test1":"","test2":"test2","test3":"test3=test3"}},"time":1476376036,"timeNano":1476376036337629959}
```

## FAQ:
### wrong Hyper.sh Credential
```
$ go run hyper-client.go --addr=147.x5.x5.x7:6443 --accessKey xxxxxxxxxxxxx --secretKey xxxxxxxxxxxxx
connecting to wss://147.x5.x5.x7:6443/events/ws
dial:websocket: bad handshake
exit status 1
```

### how to use filters query parameter in curl

REF: http://stackoverflow.com/questions/15425446/how-to-put-a-json-object-with-an-array-using-curl/26407256#26407256
```
$ curl -g 'http://127.0.0.1:2375/events?filters={"label":{"test1":true,"test2=test2","test3=test3=test3":true}}'

- "test1":true 表示存在key是test1的label
- "test1=":true 表示表示key为test1的label的value为""，精确匹配
- "test1=aaa":true 表示表示key为test1的label的value为"aaa"，精确匹配
- 多个label之间是与的关系
```


# REF:

- WebSocket:
  - http://stackoverflow.com/questions/29324251/gorilla-websocket-with-cookie-authentication
  - https://github.com/gorilla/websocket/blob/master/client_server_test.go#L322
- Watch:
  - https://github.com/kubernetes/kubernetes/blob/master/pkg/client/cache/reflector.go#L362
  - https://github.com/kubernetes/kubernetes/blob/master/pkg/apiserver/watch_test.go
- Browser aws4 example
  - https://github.com/mhart/aws4/tree/master/browser
- Parse Query Parameter - filters
  - https://github.com/docker/engine-api/blob/master/types/filters/parse_test.go
  - http://fossies.org/linux/kubernetes/pkg/registry/pod/strategy_test.go
- Event
  - https://github.com/docker/engine-api/blob/master/client/events_test.go#L56
