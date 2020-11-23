package main

import (
	"fmt"
	"sync"

	fileProto "FileService/proto"

	"golang.org/x/net/context"
)

const sessKeyLen = 10

type FileManager struct {
	mu sync.RWMutex
}

func NewFileManager() *FileManager {
	return &FileManager{
		mu: sync.RWMutex{},
	}
}

func (sm *FileManager) SetAvatar(context.Context, *fileProto.Avatar) (*fileProto.Nothing, error) {
	fmt.Printf("SEVER CALL SetAvatar")
	return &fileProto.Nothing{Dummy: true}, nil
}

func (sm *FileManager) GetAvatar(context.Context, *fileProto.User) (*fileProto.Avatar, error) {
	fmt.Printf("SEVER CALL GetAvatar")
	return &fileProto.Avatar{Email: "s@mail.ru", ImgType: "png"}, nil
}

func (sm *FileManager) SaveFile(context.Context, *fileProto.File) (*fileProto.Nothing, error) {
	fmt.Printf("SEVER CALL SaveFile")
	return &fileProto.Nothing{Dummy: true}, nil
}

func (sm *FileManager) GetFile(context.Context, *fileProto.LetterId) (*fileProto.File, error) {
	fmt.Printf("SEVER CALL GetFile")
	return &fileProto.File{LetterId: 29, FileType: "png"}, nil
}
