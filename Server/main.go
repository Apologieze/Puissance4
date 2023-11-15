package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func main() {
	var c [2]net.Conn
	var writer [2]*bufio.Writer

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
	defer c[0].Close()
	writer[0] = bufio.NewWriter(c[0])
	writer[0].WriteString("Connection au serveur réussie, tu es le joueur 1\n")
	writer[0].Flush()
	log.Println("Le client 1 s'est connecté à l'adresse", c[0].RemoteAddr().String())

	c[1], err = listener.Accept()
	if err != nil {
		log.Println("accept error:", err)
		return
	}
	defer c[1].Close()
	writer[1] = bufio.NewWriter(c[1])
	writer[1] = bufio.NewWriter(c[1])
	writer[1].WriteString("Connection au serveur réussie, tu es le joueur 2\n")
	writer[1].Flush()
	log.Println("Le client 2 s'est connecté à l'adresse", c[1].RemoteAddr().String())

	writer[0].WriteString("Le joueur 2 vient de se connecter\n")
	writer[0].Flush()

	time.Sleep(10 * time.Second)
	log.Println("Fermerture du serveur automatique")
}
