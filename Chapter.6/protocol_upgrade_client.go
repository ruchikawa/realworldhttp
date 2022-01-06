package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	// open TCP Socket
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	conn, err := dialer.Dial("tcp", "localhost:18888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// Create Request and Write to Socket directly
	request, _ := http.NewRequest("GET", "http://localhost:18888/upgrade", nil)
	request.Header.Set("Connection", "Upgra")
	request.Header.Set("Upgrade", "MyProtocol")
	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	resp, err := http.ReadResponse(reader, request)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)

	counter := 10
	for {
		data, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		fmt.Println("<-", string(bytes.TrimSpace(data)))
		fmt.Fprintf(conn, "%d\n", counter)
		fmt.Println("->", counter)
		counter--
	}
}
