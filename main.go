package main

import (
	"dino/game"
	"log"
)

func main() {
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
