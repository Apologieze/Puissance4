package main

import (
	"bufio"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("Dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("Je suis connecté")

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Read error", err)
		}
		log.Println(message)
	}
}
