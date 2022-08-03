
protoc --go_out=. --go-grpc_out=. .\api\hello.proto

go build -ldflags "-s -w" -o grpc-test.exe .\cmd\cli\main.go