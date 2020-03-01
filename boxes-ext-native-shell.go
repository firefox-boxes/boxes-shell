package main

import (
	"encoding/json"
	"fmt"
	"net/rpc"
	"strings"
	"unsafe"
)

type Msg struct {
	msg string
}

func encodeMessage(msg string) ([4]byte, []byte) {
	b, err := json.Marshal(Msg{msg:msg})
	if err != nil {
		panic(err)
	}
	bl := uint32(len(b))
	return (*[4]byte)(unsafe.Pointer(&bl))[:], b
}

func main() {
	client, _ := rpc.Dial("tcp", "127.0.0.1:6688")
	defer client.Close()
	fmt.Println("<connected>")
	for {
		cmdString := readInput()
		cmdString = strings.TrimSuffix(cmdString, "\n")
		bl, b := encodeMessage(query(client, cmdString))
		fmt.Print(bl)
		fmt.Print(b)
	}
}