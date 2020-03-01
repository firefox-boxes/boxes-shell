package main

import (
	"fmt"
	"net/rpc"
	"strings"
)

func main() {
	client, _ := rpc.Dial("tcp", "127.0.0.1:6688")
	defer client.Close()
	fmt.Println("<connected>")
	for {
		cmdString := readInput()
		cmdString = strings.TrimSuffix(cmdString, "\n")
		fmt.Println(query(client, cmdString))
	}
}