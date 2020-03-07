package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/rpc"
	"os"
	"strconv"
	"strings"
	"unsafe"

	"github.com/firefox-boxes/boxes-shell/logging"
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
	logging.Info("Sending `%v`", msg)
	bl, b := encodeMessage(msg)
	fmt.Print(string(bl))
	fmt.Print(string(b))
}

func decode(msg []byte) string {
	var command string
	json.Unmarshal(msg, &command)
	logging.Info("Arrival:3 m=%v decoded=%v", msg, command)
	return command
}

func readInput() []byte {
	reader := bufio.NewReader(os.Stdin)
	length := [4]byte{}
	n, err := io.ReadFull(reader, length[:])
	if err != nil {
		logging.Info("%v bytes read: %v", err, n)
		panic(err)
	}
	l := *(*uint32)(unsafe.Pointer(&length))
	logging.Info("Arrival:1 lb=%v l=%v", length, l)
	input := make([]byte, l, l)
	n, err = io.ReadFull(reader, input)
	if err != nil {
		logging.Info("%v bytes read: %v", err, n)
		panic(err)
	}
	logging.Info("Arrival:2 content=%v", logging.ProcessStr(string(input), 40))
	return input
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:6688")
	if err != nil {
		logging.Info(err.Error())
		panic(err)
	}
	logging.Info("Connected to 127.0.0.1:6688")
	defer client.Close()
	sendMessage(query(client, "whoami " + strconv.Itoa(os.Getppid())))
	for {
		cmdString := decode(readInput())
		cmdString = strings.TrimSuffix(cmdString, "\n")
		sendMessage(query(client, cmdString))
	}
}