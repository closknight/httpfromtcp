package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening tcp traffic %s", err.Error())
	}
	defer listener.Close()

	fmt.Println("Listening for traffic on", port)
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Println("accepted connection from", connection.RemoteAddr())
		linesChan := getLinesChannel(connection)
		for line := range linesChan {
			fmt.Println(line)
		}
		fmt.Println("connection from ", connection.RemoteAddr(), "closed")
	}

	// for line := range getLinesChannel() {
	// 	fmt.Printf("read: %s\n", line)
	// }
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string)
	go func() {
		defer f.Close()
		defer close(out)
		currentLine := ""
		for {
			buffer := make([]byte, 8, 8)
			n, err := f.Read(buffer)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}
			str := string(buffer[:n])
			parts := strings.Split(str, "\n")
			for i := range len(parts) - 1 {
				out <- currentLine + parts[i]
				currentLine = ""
			}
			currentLine += parts[len(parts)-1]
		}
		out <- currentLine
	}()
	return out
}
