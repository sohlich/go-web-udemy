package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("Error occured during listener creation" + err.Error())
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		log.Println("Received connection")
		if err != nil {
			panic("Error occured")
		}
		log.Println("Connection going to be handled")
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	requestLine := ""
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if requestLine == "" {
			requestLine = scanner.Text()
		}
		if scanner.Text() == "" {
			mux(conn, requestLine)
			break
		}
	}
}

type Handler func(net.Conn)

func mux(conn net.Conn, head string) {
	header := strings.Split(head, " ")
	var handler Handler
	log.Printf("%s is being handled.\n")
	switch header[0] {
	case "GET":
		handler = handleGet
	case "POST":
		handler = handlePost
	default:
		return
	}
	handler(conn)
}

var response = `
<html>\r\n
<body>\r\n
<h1>Hello, World from %s method.</h1>\r\n
</body>\r\n
</html>\r\n`

func handleGet(conn net.Conn) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\n")
	fmt.Fprintf(conn, "Content-Length: %d\n", len(response))
	fmt.Fprint(conn, "Content-Type: text/html\n")
	fmt.Fprint(conn, "\n")
	fmt.Fprint(conn, fmt.Sprintf(response, "GET"))
}

func handlePost(conn net.Conn) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\n")
	fmt.Fprintf(conn, "Content-Length: %d\n", len(response))
	fmt.Fprint(conn, "Content-Type: text/html\n")
	fmt.Fprint(conn, "\n")
	fmt.Fprint(conn, fmt.Sprintf(response, "POST"))
}
