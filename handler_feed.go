package main

import (
	"context"
	"fmt"
	"github.com/atulantonyz/gator/internal/database"
	"github.com/google/uuid"
	"os"
	"time"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		fmt.Printf("usage: %s <name> <url>\n", cmd.Name)
		os.Exit(1)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return err
	}
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}
	fmt.Println("Feed was created")
	fmt.Printf("id: %v\ncreated_at: %v\nupdated_at: %v\nname: %s\nurl: %s\nuser_id: %v\n", feed.ID, feed.CreatedAt, feed.UpdatedAt,
		feed.Name, feed.Url, feed.UserID)
	return nil

}

func handlerFeeds(s *state, cmd command) error {
	feedsRows, err := s.db.GetFeeds(context.Background())
	if err != nil {
		fmt.Println("Unable to retrieve feeds")
		os.Exit(1)
	}
	for _, feed := range feedsRows {
		fmt.Printf("name: %s\nurl: %s\nuser: %s\n", feed.Name, feed.Url, feed.Name_2)
	}

	return nil
}
