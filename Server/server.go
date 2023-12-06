package main

import (
	"log"
	"net"
)

var c [2]net.Conn
var buffer [2][]byte
var state int = titleState

func starting_server() {
	buffer[0] = make([]byte, 128)
	buffer[1] = make([]byte, 128)

	log.Println("Attente de connection des joueurs...")

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Println("listen error:", err)
		return
	}
	defer listener.Close()

	c[0], err = listener.Accept()
	if err != nil {
		log.Println("accept error:", err)
		return
	}
	//defer c[0].Close()

	_, err = c[0].Write([]byte{0, 1})
	log.Println("Le client 1 s'est connecté à l'adresse", c[0].RemoteAddr().String())

	c[1], err = listener.Accept()
	if err != nil {
		log.Println("accept error:", err)
		return
	}
	//defer c[1].Close()
	_, err = c[1].Write([]byte{0, 2})
	_, err = c[0].Write([]byte{1})
	_, err = c[1].Write([]byte{1})
	log.Println("Le client 2 s'est connecté à l'adresse", c[1].RemoteAddr().String())
	state = colorSelectState
}

func main() {
	starting_server()

	for {
		update()
	}
}
