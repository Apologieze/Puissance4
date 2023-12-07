package main

import (
	"log"
	"net"
)

type connection struct {
	g         *game
	connected bool
	conn      net.Conn
	sending   chan []byte
}

func (c *connection) startingConnection() {
	c.sending = make(chan []byte, 1)
	go c.sendingServer()
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
			c.g.playerId = buffer[1]
			c.g.turn = int(c.g.playerId) - 1
			log.Println("Je suis le joueur", c.g.playerId)
			if c.g.playerId == 2 {
				c.connected = true
				log.Println("Tous les joueurs sont connectés")
			}
		case 1:
			switch c.g.gameState {
			case titleState:
				c.connected = true
				if c.g.playerId == 1 {
					log.Println("Le joueur 2 c'est connecté")
				}
				log.Println("Tous les joueurs sont connectés")

			case colorSelectState:
				log.Println("Tous les joueurs ont choisi leur couleur")
				c.g.gameState++
			}
		case 2:
			c.g.opponentLastPos = int(buffer[1])
		}
	}
}

func (c *connection) sendingServer() {
	for {
		_, err := c.conn.Write(<-c.sending)
		if err != nil {
			log.Println("accept error:", err)
			return
		}
	}
}
