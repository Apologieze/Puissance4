package main

import (
	"fmt"
	"log"
)

func updatePlayer(id int, ch chan []byte) {
	for {
		_, err := c[id].Read(buffer[id])
		if err != nil {
			log.Fatal("Read error", err)
		}
		ch <- buffer[id]
	}
}

func update() {
	switch state {
	case colorSelectState:
		select {
		case msg := <-ch1:
			if msg[0] == 2 {
				selectedP1Color = int(msg[1])
				fmt.Println(selectedP1Color)
				if selectedP2Color != -1 {
					state++
				} else {
					c[1].Write([]byte{4, byte(selectedP1Color)})
				}
			} else if msg[0] == 3 {
				c[1].Write(msg)
			}
		case msg := <-ch2:
			if msg[0] == 2 {
				selectedP2Color = int(msg[1])
				fmt.Println(selectedP2Color)
				if selectedP1Color != -1 {
					state++
				} else {
					c[0].Write([]byte{4, byte(selectedP2Color)})
				}
			} else if msg[0] == 3 {
				c[0].Write(msg)
			}
		}
	case startingState:
		_, err := c[0].Write([]byte{1, byte(selectedP2Color)})
		_, err = c[1].Write([]byte{1, byte(selectedP1Color)})
		if err != nil {
			log.Fatal("Read error", err)
		}
		state++

	case playState:
		select {
		case msg := <-ch1:
			c[1].Write([]byte{2, msg[1]})
		case msg := <-ch2:
			c[0].Write([]byte{2, msg[1]})
		}
	}
}
