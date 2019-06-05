package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

type ClientJob struct {
	message string
	conn    net.Conn
}

func main() {
	// connect to this socket
	conn, _ := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)

	defer conn.Close()

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')

		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err == nil {
			fmt.Print("Message from server: " + message)
		} else {
			fmt.Print(err)
		}

	}
}
