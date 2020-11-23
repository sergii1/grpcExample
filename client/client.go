package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	fileProto "FileService/proto"
)

func main() {

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	fileManager := fileProto.NewFileServiceClient(grcpConn)

	ctx := context.Background()

	avatar, _ := fileManager.GetAvatar(ctx, &fileProto.User{Email: "suko"})
	fmt.Println(avatar.Email)
}