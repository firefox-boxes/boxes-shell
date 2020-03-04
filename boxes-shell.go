package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return str
}

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