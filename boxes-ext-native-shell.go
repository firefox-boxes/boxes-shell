package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/rpc"
	"os"
	"strconv"
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

func decode(msg string) string {
	var command string
	json.Unmarshal([]byte(msg), &command)
	return command
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	length := make([]byte, 4, 4)
	reader.Read(length)
	l := *(*uint32)(unsafe.Pointer(&length))
	input := make([]byte, l, l)
	reader.Read(input)
	return string(input)
}

func main() {
	client, _ := rpc.Dial("tcp", "127.0.0.1:6688")
	defer client.Close()
	sendMessage(query(client, "whoami " + strconv.Itoa(os.Getppid())))
	for {
		cmdString := decode(readInput())
		cmdString = strings.TrimSuffix(cmdString, "\n")
		sendMessage(query(client, cmdString))
	}
}