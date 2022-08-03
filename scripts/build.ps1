
protoc --go_out=. --go-grpc_out=. .\api\hello.proto

go build -ldflags "-s -w" -o hello-client.exe .\cmd\hello-client\main.go
go build -ldflags "-s -w" -o hello-local.exe .\cmd\hello-local\main.go
go build -ldflags "-s -w" -o hello-server.exe .\cmd\hello-server\main.go
