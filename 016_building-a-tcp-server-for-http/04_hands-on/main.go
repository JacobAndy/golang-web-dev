package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}
func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)
			fmt.Println(m)
			t := m[0]
			u := m[1]
			if t == "GET" {
				if u == "/" {
					response(conn, `<!DOCTYPE html><html><head><meta charset="UTF-8">
					<title>Hello World</title></head><body><h1>GET method to /</h1></body></html>`)
				} else if u == "/info" {
					response(conn, `<!DOCTYPE html><html><head><meta charset="UTF-8">
					<title>Hello World</title></head><body><h1>GET method to /info</h1></body></html>`)
				}
			} else if t == "PUT" {
				if u == "/" {
					response(conn, `<!DOCTYPE html><html><head><meta charset="UTF-8">
					<title>Hello World</title></head><body><h1>PUT method to /</h1></body></html>`)
				} else if u == "/api" {
					response(conn, `<!DOCTYPE html><html><head><meta charset="UTF-8">
					<title>Hello World</title></head><body><h1>PUT method to /api</h1></body></html>`)
				}
			}
		}
		if ln == "" {
			break
		}
		i++
	}
}

func response(conn net.Conn, html string) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, html)
}
