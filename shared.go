package main

import (
	"net/rpc"
)

func query(client *rpc.Client, req string) string {
	var res string
	client.Call("IPC.Handle", &req, &res)
	return res
}