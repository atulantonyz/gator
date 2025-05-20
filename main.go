package main

import (
	"github.com/atulantonyz/gator/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	var app_state state
	app_state.cfg = &cfg
	app_commands := commands{
		map[string]func(*state, command) error{},
	}
	app_commands.register("login", handlerLogin)
	cmd_line_args := os.Args
	if len(cmd_line_args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}

	cmdName := cmd_line_args[1]
	cmdArgs := cmd_line_args[2:]

	err = app_commands.run(&app_state, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}

}
