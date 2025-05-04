package main

import (
	"fmt"
	"github.com/atulantonyz/gator/internal/config"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read config: %v\n", cfg)
	cfg.SetUser("Atul")
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read config again: %v\n", cfg)

}
