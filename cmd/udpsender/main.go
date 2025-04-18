package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const address = "localhost:42069"

func main() {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalf("could not resove address %s: %s", address, err.Error())
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("error connecting to udp: %s", err.Error())
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error reading input: %s\n", err.Error())
		} else {
			_, err := conn.Write([]byte(line))
			if err != nil {
				fmt.Printf("error writing to connection: %s\n", err.Error())
			}
		}
	}
}
