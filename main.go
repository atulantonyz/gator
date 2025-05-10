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

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Username is required")
	}
	s.cfg.SetUser(cmd.args[0])
	fmt.Println("User has been set to " + cmd.args[0])
	return nil
}

type commands struct {
	cmd_list map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.cmd_list[cmd.name]
	if !ok {
		return fmt.Errorf("Command does not exist")
	}
	err := handler(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmd_list[name] = f
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
