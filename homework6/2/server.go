package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	port := 8081
	fmt.Printf("Server starting on port: %d \nNow run the client.", port)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	onErrPanic(err)

	conn, err := ln.Accept()
	onErrPanic(err)

	for {
		inputMsg, err := bufio.NewReader(conn).ReadString('\n')
		onErrPanic(err)

		inputMsg = strings.Trim(inputMsg, "\n")

		fmt.Printf("Message received: %v\n", inputMsg)

		check, err := strconv.Atoi(inputMsg)

		var outputMsg string

		if err != nil {
			outputMsg = strings.ToUpper(inputMsg)
		} else {
			outputMsg = strconv.Itoa(check * 2)
		}

		conn.Write([]byte(outputMsg + "\n"))
	}
}

func onErrPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}
