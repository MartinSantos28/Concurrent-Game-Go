package main

import (
	"dino/scene"
	"log"
)

func main() {
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
