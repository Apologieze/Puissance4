package main

import (
	"log"
	"net"
)

type connection struct {
	g         *game
	playerId  byte
	connected bool
	conn      net.Conn
}

func (c *connection) startingConnection() {
	var err error
	c.conn, err = net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("Dial error:", err)
		return
	}
	defer c.conn.Close()
	log.Println("Je suis connecté")
	buffer := make([]byte, 128)

	for {
		_, err = c.conn.Read(buffer)
		if err != nil {
			log.Fatal("Read error", err)
		}
		switch buffer[0] {
		case 0:
			c.playerId = buffer[1]
			log.Println("Je suis le joueur", c.playerId)
		case 1:
			switch c.g.gameState {
			case titleState:
				c.connected = true
				if c.playerId == 1 {
					log.Println("Le joueur 2 c'est connecté")
				}
				log.Println("Tous les joueurs sont connectés")

			case colorSelectState:
				log.Println("Tous les joueurs ont choisi leur couleur")
				c.g.gameState++
			}
		}
	}
}

func (c *connection) sendingServer(data []byte) {
	_, err := c.conn.Write(data)
	if err != nil {
		log.Println("accept error:", err)
		return
	}
}
