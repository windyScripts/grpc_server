protoc --go_out=. --go-grpc_out=. proto/main.proto

go mod init <name>

go get google.golang.org/grpc

For generating from openssl.cnf file:

openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem -config cert.conf