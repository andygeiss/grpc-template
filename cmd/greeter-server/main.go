package main

import (
	"flag"
	"log"
	"net"

	"github.com/andygeiss/grpc-template/internal/greeter"
	"google.golang.org/grpc"
)

func main() {
	// parse the flags
	listen := flag.String("listen", ":3000", "")
	flag.Parse()
	// create a new listening socket
	lis, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	// create a gRPC server and register the master
	grpcServer := grpc.NewServer()
	service := greeter.NewGreeter()
	greeter.RegisterGreeterServer(grpcServer, service)
	log.Printf("start listening at %s ...", *listen)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
