Generate Certificate:
  - generate private key
  openssl genrsa -out `server.key` 2048
  - generate public key
  openssl rsa -in server.key -out `server.key.public`
  - genereate certificate
  openssl req -new -x509 -key server.key -out `server.crt` -days 365

Start server:
  go run server.go

Start client
  go run client.go

Open Web Browser:
  https://$:8888
