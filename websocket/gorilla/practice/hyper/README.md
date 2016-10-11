websocket client for hyper
==================================================

>Connect to websocket api `/events/ws` of Hyper.sh apirouter

- SSL: `InsecureSkipVerify: true`
- Sign4: Use util/sign4.go

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

# Usage
```
//server
https://github.com/getdvm/hyper-api-router/pull/298

//client
$ go get github.com/Jimmy-Xu/learn-go/websocket/gorilla/practice/hyper
$ cd $GOPATH/src/github.com/Jimmy-Xu/learn-go/websocket/gorilla/practice/hyper
$ go run hyper-client.go --accessKey KXARxxxxxxxxxxxxx5WR8 --secretKey Ema5xxxxxxxxxxxxxxxxxxxxxxxxxxxZgNe
connecting to wss://147.75.195.37:6443/events/ws
recv[json]: {
  "Action": "ADDED",
  "Actor": {
    "Attributes": {
      "ExitCode": "",
      "Image": "busybox",
      "Labels": null,
      "Name": "test2"
    },
    "Id": "d180511d5b5c9e34c769adbf6a4c292249a26e71fc64e54aef67fcc0940dad45"
  },
  "From": "busybox",
  "Id": "d180511d5b5c9e34c769adbf6a4c292249a26e71fc64e54aef67fcc0940dad45",
  "Status": "Running",
  "Time": 1.47601746e+09,
  "TimeNano": 1.47601746053364e+18,
  "Type": "container"
}
```

## FAQ:
### wrong Hyper.sh Credential
```
$ go run hyper-client.go --accessKey xxxxxxxxxxxxx --secretKey xxxxxxxxxxxxx
connecting to wss://147.75.195.37:6443/events/ws
dial:websocket: bad handshake
exit status 1
```
