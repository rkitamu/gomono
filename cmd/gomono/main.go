package main

import (
	"log"

	"github.com/rkitamu/gomono/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
