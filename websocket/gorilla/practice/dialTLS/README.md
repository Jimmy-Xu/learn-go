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
