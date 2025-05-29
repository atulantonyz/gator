package main

import (
	"context"
	"fmt"
	"github.com/atulantonyz/gator/internal/database"
	"github.com/google/uuid"
	"os"
	"time"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		fmt.Println("User does not exist, cannot login!")
		os.Exit(1)
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User has been set to " + cmd.Args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}
	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		fmt.Println("User with given name already exists")
		os.Exit(1)
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Println("User was created")
	fmt.Printf("id: %v\ncreated_at: %v\nupdated_at: %v\nname: %s\n", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Println("Unable to reset users table")
		os.Exit(1)
	}
	fmt.Println("Users table successfully reset")
	return err
}
