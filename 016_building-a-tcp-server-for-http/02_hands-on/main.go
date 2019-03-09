package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	t, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer t.Close()
	for {
		conn, err := t.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)

	respond(conn)
}
func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)
			fmt.Println("URL on request: ", m[1])
		}
		if ln == "" {
			break
		}
		i++
	}
}
func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html><head><meta charset="UTF-8">
	<title>Hello World</title></head><body><h1>Hello World</h1></body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
