package main

import (
	"context"
	"fmt"
	"github.com/atulantonyz/gator/internal/database"
	"github.com/google/uuid"
	"os"
	"strconv"
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

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("Unable to retrieve users")
		os.Exit(1)
	}
	for _, user := range users {
		fmt.Printf("* %s", user)
		if s.cfg.Current_user_name == user {
			fmt.Printf(" (current)")
		}
		fmt.Printf("\n")
	}
	return nil
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s <name> (limit)", cmd.Name)
	}
	limit := "2"

	if len(cmd.Args) == 1 {
		limit = cmd.Args[0]
	}
	limit_int, err := strconv.ParseInt(limit, 10, 32)
	if err != nil {
		return err
	}
	posts, err := s.db.GetPostsForUser(
		context.Background(), database.GetPostsForUserParams{
			UserID: user.ID,
			Limit:  int32(limit_int),
		},
	)
	for _, post := range posts {
		fmt.Printf("\n* %s : %s\n", post.Title, post.Url)
		fmt.Printf("- published: %s\n", post.PublishedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("\n %s\n", post.Description.String)
	}

	return nil
}
