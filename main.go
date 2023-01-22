package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	server, err := net.Listen("tcp", "localhost:7071")

	checkErr(err)

	for {
		conn, err := server.Accept()

		checkErr(err)

		conn.Read()

	}

}

func checkErr(err error) {

	if err != nil {

		log.Fatal(err)
	}
}

func dialTcp() {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')

}
