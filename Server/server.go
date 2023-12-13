package main

import (
	"log"
	"net"
)

const (
	LOCAL = "localhost:8081"
	VM    = ":80"
)

var c [2]net.Conn
var buffer [2][]byte
var state int = titleState
var ch1 chan []byte
var ch2 chan []byte
var selectedP1Color int
var selectedP2Color int

func starting_server() {
	buffer[0] = make([]byte, 128)
	buffer[1] = make([]byte, 128)

	ch1 = make(chan []byte, 1)
	ch2 = make(chan []byte, 1)

	selectedP1Color, selectedP2Color = -1, -1

	log.Println("Attente de connection des joueurs...")

	listener, err := net.Listen("tcp", LOCAL)
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
	//_, err = c[1].Write([]byte{1})
	log.Println("Le client 2 s'est connecté à l'adresse", c[1].RemoteAddr().String())
	state = colorSelectState
}

func main() {
	starting_server()

	go updatePlayer(0, ch1)
	go updatePlayer(1, ch2)
	for {
		update()
	}
}
