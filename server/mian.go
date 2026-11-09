package main

import (
	"grpcEx/pkg"
	"grpcEx/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	srv := &pkg.Server{}

	grpcSerever := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(grpcSerever, srv)
	if err := grpcSerever.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
