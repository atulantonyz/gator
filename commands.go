package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmd_list map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.cmd_list[cmd.name]
	if !ok {
		return fmt.Errorf("Command does not exist")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmd_list[name] = f
}
