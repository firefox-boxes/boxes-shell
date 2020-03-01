package main

import (
	"fmt"
	"net/rpc"
	"strings"
)

func main() {
	client, _ := rpc.Dial("tcp", "127.0.0.1:6688")
	defer client.Close()
	fmt.Println("boxes-ipc at 127.0.0.1:6688")
	for {
		fmt.Print(">>> ")
		cmdString := readInput()
		cmdString = strings.TrimSuffix(cmdString, "\n")
		fmt.Println(query(client, cmdString))
		fmt.Println()
	}
}