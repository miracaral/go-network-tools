package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err != nil {
		log.Fatalln("Unable to connect to destination")

	}
	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Unable to port ")

	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
