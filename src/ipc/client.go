package ipc

import (
	"encoding/json"
)

type IpcClient struct {
	conn chan string
}

func NewIPcClient(server* IpcServer) *IpcClient {
	
}