package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting TCP server on localhost:8080")

	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error Listening :", err.Error())

		os.Exit(1)
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error Connecting:", err.Error())
			return
		}

		fmt.Println("Client " + c.LocalAddr().String() + " connected.")

		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	for {
		buffer, err := bufio.NewReader(c).ReadBytes('\n')
		if err != nil {
			log.Println("Client " + c.LocalAddr().String() + " left")
			return
		}

		log.Println("Client Message:", string(buffer[:len(buffer)-1]))

		c.Write(buffer)
	}
}
