package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Println("Unable to reset users table")
		os.Exit(1)
	}
	fmt.Println("Users table successfully reset")
	return err
}
