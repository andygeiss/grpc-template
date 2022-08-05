
protoc --go_out=. --go-grpc_out=. .\api\greeter.proto

go build -ldflags "-s -w" -o greeter-client.exe .\cmd\greeter-client\main.go
go build -ldflags "-s -w" -o greeter-local.exe .\cmd\greeter-local\main.go
go build -ldflags "-s -w" -o greeter-server.exe .\cmd\greeter-server\main.go
