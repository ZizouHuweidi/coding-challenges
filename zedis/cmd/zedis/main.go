package main

import (
	"io"
	"log"
	"net"
)

func main() {
	log.Println("Listening on port :6379")

	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("error reading from client", err.Error())
		}
	}

	conn.Write([]byte("+OK\r\n"))
}
