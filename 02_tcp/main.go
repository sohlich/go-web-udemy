package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
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
	go handle(conn)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
	fmt.Println("Captured %v. Exiting...")
	os.Exit(0)
}

func handle(conn net.Conn) {
	log.Println("Connection being handled")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		io.Copy(conn, strings.NewReader("OK"))
	}
}
