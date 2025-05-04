package main

import (
	"fmt"
	"github.com/atulantonyz/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error in reading config")
	}
	cfg.SetUser("Atul")
	cfg2, err := config.Read()
	if err != nil {
		fmt.Println("Error in reading config")
	}
	fmt.Println(cfg2)

}
