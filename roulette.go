package main

import (
	"net"
	"log"
	"fmt"
	"io"
)

var friendQueue = make(chan net.Conn)

func main() {
	service := ":8000"
	l, err := net.Listen("tcp", service)
	if err != nil {
		log.Fatal(err)
	}

	go match()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Fprintf(conn, "Waiting other friend...\n")
		friendQueue <- conn
	}
}

func match() {
	for {
		c1 := <-friendQueue
		c2 := <-friendQueue

		fmt.Fprintf(c1, "Matched a new friend. Say Hello...\n")
		fmt.Fprintf(c2, "Matched a new friend. Say Hello...\n")


		go letChat(c1, c2)
		go letChat(c2, c1)
	}
}

func letChat(c1, c2 net.Conn) {
	defer c1.Close()
	defer c2.Close()

	_, err := io.Copy(c1, c2)
	if err != nil {
		log.Println(err)
		return
	}
}
