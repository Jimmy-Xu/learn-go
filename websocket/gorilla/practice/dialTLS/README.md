# Certificate without CA

Generate Certificate for server:
  - generate private key
  openssl genrsa -out server.key 2048
  - generate public key
  openssl rsa -in server.key -out server.key.public
  - generate certificate
  openssl req -new -x509 -key server.key -out server.crt -days 365



# Certificate with CA

Generate CA Certificate:
  - genreate ca key
  openssl genrsa -out ca.key 2048
  - genreate ca cert
  openssl req -x509 -new -nodes -key ca.key -subj "/CN=localhost" -days 5000 -out ca.crt

Generate Certificate for server:
  - genreate server key
  openssl genrsa -out server.key 2048
  - generate CSR (Certificate Signing Request)
  openssl req -new -key server.key -subj "/CN=localhost" -out server.csr
  - generate server cert
  openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000

Generate Certificate for client:
  - genreate client key
  openssl genrsa -out client.key 2048
  - generate CSR (Certificate Signing Request)
  openssl req -new -key client.key -subj "/CN=localhost" -out client.csr
  - generate client cert
  openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 5000



# Run

Start server:
  go run server.go

Start client
  go run client.go

Open Web Browser:
  https://$:8888

==================================================

- SSL
- Authentication
- Sign4


http://stackoverflow.com/questions/29324251/gorilla-websocket-with-cookie-authentication
https://github.com/gorilla/websocket/blob/master/client_server_test.go#L322

===================================================

# hyper-client.go

>Connect to websocket api `/events/ws` of Hyper.sh apirouter

- `InsecureSkipVerify: true`
- Use util/sign4.go

## Run result
```
//client
$ go run hyper-client.go
connecting to wss://147.75.195.37:6443/events/ws
URL:wss://147.75.195.37:6443/events/ws
recv: 2016-10-08 14:16:27.183389957 +0800 CST
recv: 2016-10-08 14:16:28.183427987 +0800 CST
recv: 2016-10-08 14:16:29.183434709 +0800 CST

//apirouter: https://github.com/getdvm/hyper-api-router/pull/298
I1008 06:16:25.812586   28249 key.go:31] check V4 signature
I1008 06:16:25.813984   28249 tbac.go:122] Authorize tenant 1ca724a038554e89a2647100b004d2fc (restricted: false) via key jimmy+test@hyper.sh for action , status 0
I1008 06:16:25.813996   28249 tbac.go:72] Authorize for resource  from user jimmy+test@hyper.sh, result : <nil>
I1008 06:16:25.814085   28249 log.go:8] APIRequest[1475907385814082037]: Start, Tenant: 1ca724a038554e89a2647100b004d2fc, Verb: GET, Path: /events/ws
```

## FAQ:
### wrong Hyper.sh Credential
```
$ go run hyper-client.go
connecting to wss://147.75.195.37:6443/events/ws
URL:wss://147.75.195.37:6443/events/ws
dial:websocket: bad handshake
exit status 1
```
