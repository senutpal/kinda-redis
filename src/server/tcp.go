package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/senutpal/kinda-redis/src/resp"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	fmt.Println("Server started on port 6379")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return
	}
	_, parsedResp := resp.ReadNextRESP(buf[:n])

	var args []string

	if parsedResp.Type == resp.Array {
		parsedResp.ForEach(func(r resp.RESP) bool {
			args = append(args, r.String())
			return true
		})
	}

	if len(args) > 0 {
		command := strings.ToUpper(args[0])

		if command == "PING" {
			conn.Write([]byte("+PONG\r\n"))
		} else if command == "ECHO" && len(args) > 1 {
			wordToEcho := args[1]

			formattedResponse := resp.AppendBulkString(nil, wordToEcho)
			conn.Write(formattedResponse)
		}
	}
}
