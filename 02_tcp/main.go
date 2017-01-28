package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("Error occured during listener creation" + err.Error())
	}
	conn, err := listener.Accept()
	log.Println("Received connection")
	if err != nil {
		panic("Error occured")
	}
	log.Println("Connection going to be handled")
	handle(conn)
}

func handle(conn net.Conn) {
	log.Println("Connection being handled")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		io.Copy(conn, strings.NewReader("OK"))
	}
}
