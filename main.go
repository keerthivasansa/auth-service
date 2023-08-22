package main

import (
	"kv/auth/auth_grpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()

	auth_grpc.RegisterAuthGrpc(grpcServer)

	log.Fatal(grpcServer.Serve(ls))
}
