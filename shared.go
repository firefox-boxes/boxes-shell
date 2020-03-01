package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
)

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return str
}

func query(client *rpc.Client, req string) string {
	var res string
	client.Call("IPC.Handle", &req, &res)
	return res
}