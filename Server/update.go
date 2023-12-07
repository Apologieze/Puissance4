package main

import (
	"fmt"
	"log"
)

func update() {
	switch state {
	case colorSelectState:
		fmt.Println("color")
		_, err := c[0].Read(buffer[0])
		if err != nil {
			log.Fatal("Read error", err)
		}
		fmt.Println(buffer[0][1])

		_, err = c[1].Read(buffer[1])
		if err != nil {
			log.Fatal("Read error", err)
		}
		fmt.Println(buffer[1][1])

		_, err = c[0].Write([]byte{1})
		_, err = c[1].Write([]byte{1})
		state++

	case playState:
		_, err := c[0].Read(buffer[0])
		if err != nil {
			log.Fatal("Read error", err)
		}
		fmt.Println(buffer[0][1])
		_, err = c[1].Write([]byte{2, buffer[0][1]})

		_, err = c[1].Read(buffer[1])
		if err != nil {
			log.Fatal("Read error", err)
		}
		fmt.Println(buffer[1][1])
		_, err = c[0].Write([]byte{2, buffer[1][1]})
	}
}
