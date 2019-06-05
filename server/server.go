package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

type ClientJob struct {
	message string
	conn    net.Conn
}

func generateResponses(clientJobs chan ClientJob) {
	for {
		// Wait for the next job to come off the queue.
		clientJob := <-clientJobs

		// Do something thats keeps the CPU buys for a whole second.
		for start := time.Now(); time.Now().Sub(start) < time.Second; {
		}

		fmt.Printf("Recive :" + clientJob.message + ".\n")
		// Send back the response.
		clientJob.conn.Write([]byte(clientJob.message))
	}
}

func main() {

	//var master *(chan ClientJob) = nil

	clientJobs := make(chan ClientJob)
	go generateResponses(clientJobs)

	// if master == nil {
	// 	master = &clientJobs
	// }

	ln, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	check(err, "Server is ready.")

	defer ln.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := ln.Accept()
		check(err, "Accepted connection.")

		go func() {
			buf := bufio.NewReader(conn)

			for {
				message, err := buf.ReadString('\n')

				if err != nil {
					fmt.Printf("Client disconnected.\n")
					break
				}

				clientJobs <- ClientJob{message, conn}

				// if master != nil {
				// 	*master <- ClientJob{message, conn}
				// }
			}
		}()
	}

}
