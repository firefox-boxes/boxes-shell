package main

import (
	"encoding/json"
	"fmt"
	"net/rpc"
	"strings"
	"unsafe"
)

type Msg struct {
	Msg string `json:"msg"`
}

func encodeMessage(msg string) ([]byte, []byte) {
	b, err := json.Marshal(Msg{Msg:msg})
	if err != nil {
		panic(err)
	}
	bl := uint32(len(b))
	return (*[4]byte)(unsafe.Pointer(&bl))[:], b
}

func sendMessage(msg string) {
	bl, b := encodeMessage(msg)
	fmt.Print(string(bl))
	fmt.Print(string(b))
}

func main() {
	client, _ := rpc.Dial("tcp", "127.0.0.1:6688")
	defer client.Close()
	sendMessage("<connected>")
	for {
		cmdString := readInput()
		cmdString = strings.TrimSuffix(cmdString, "\n")
		sendMessage(query(client, cmdString))
	}
}