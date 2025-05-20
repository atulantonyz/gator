package main

import (
	"fmt"
	"github.com/atulantonyz/gator/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Username is required")
	}
	s.cfg.SetUser(cmd.args[0])
	fmt.Println("User has been set to " + cmd.args[0])
	return nil
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
		log.Fatalf("Too few arguments")
	}
	cmd1 := command{
		cmd_line_args[1],
		cmd_line_args[2:],
	}
	err = app_commands.run(&app_state, cmd1)
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}

}
