package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":8081")
	onErrPanic(err)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Server hears: ")
		outputMsg, err := reader.ReadString('\n')
		onErrPanic(err)

		if strings.Trim(outputMsg, "\n") == "exit" {
			os.Exit(0)
		}

		fmt.Fprintf(conn, outputMsg+"\n")
		inputMsg, err := bufio.NewReader(conn).ReadString('\n')
		onErrPanic(err)
		fmt.Printf("Server says: %s", string(inputMsg))
	}
}

func onErrPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}
