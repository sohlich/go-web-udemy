package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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

var response = `
<html>\r\n
<body>\r\n
<h1>Hello, World!</h1>\r\n
</body>\r\n
</html>\r\n`

func handle(conn net.Conn) {
	defer conn.Close()
	log.Println("Connection being handled")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if scanner.Text() == "" {
			break
		}
	}
	fmt.Fprint(conn, "HTTP/1.1 200 OK\n")
	fmt.Fprintf(conn, "Content-Length: %d\n", len(response))
	fmt.Fprint(conn, "Content-Type: text/html\n")
	fmt.Fprint(conn, "\n")
	fmt.Fprint(conn, response)

}
