package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":7878")

	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			print("Ruh roh 2")
		}

		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {

	reader := bufio.NewReader(conn)

	var err error = nil

	// consecutive_CRLF_counter := 0

	var headers map[string]string

	headers = make(map[string]string)

	var request_URL string
	var HTTP_verb string
	// var HTTP_version string

	line_number := 0
	for err == nil {
		status, err := reader.ReadString('\n')

		if err != nil {
			break
		}
		if status == "\r\n" {
			// Done with headers
			break
		}
		if line_number == 0 {
			// Request-Line
			split_line := strings.Split(status, " ")
			HTTP_verb = split_line[0]
			request_URL = split_line[1]
			// HTTP_version := split_line[2]
		}
		if strings.Contains(status, ":") {
			// Is formatted like a header

			key_value_pair := strings.SplitN(status, ":", 2)

			key := key_value_pair[0]
			val := key_value_pair[1]

			headers[key] = val

		}

		line_number++
	}
	if HTTP_verb == "GET" {
		str := "You requested " + request_URL + "!"
		println(str)
		serveGetRequest(conn, request_URL)
	}

	conn.Close()

}

func serveGetRequest(conn net.Conn, request_URL string) {
	cwd, _ := os.Getwd()
	file_path := path.Join(cwd, request_URL)
	data, _ := ioutil.ReadFile(file_path) // Catch error
	data_string := string(data)
	fmt.Fprintln(conn, data_string)
}

// Abstract request parsing,
// make a struct to hold output of that parsing
// Make a method to serve file or directory
