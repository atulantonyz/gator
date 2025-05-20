package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		fmt.Errorf("usage: %s <name>", cmd.name)
	}

	name := cmd.args[0]

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User has been set to " + cmd.args[0])
	return nil
}
