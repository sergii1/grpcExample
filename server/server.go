package main

import (
	fileProto "FileService/proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("cant listet port", err)
	}

	server := grpc.NewServer()

	fileProto.RegisterFileServiceServer(server, NewFileManager())

	fmt.Println("starting server at :8081")
	server.Serve(lis)
}