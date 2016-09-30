
Instlal:
  go get github.com/golang/net
  mkdir -p $GOPATH/src/golang.org/x
  ln -s $GOPATH/src/github.com/golang/net $GOPATH/src/golang.org/x/net
  go install golang.org/x/net/websocket


Usage:
  import "golang.org/x/net/websocket"

Doc:
  https://godoc.org/golang.org/x/net/websocket

Example:
  https://github.com/golang-samples/websocket
