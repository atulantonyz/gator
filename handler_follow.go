package main

import (
	"context"
	"fmt"
	"github.com/atulantonyz/gator/internal/database"
	"github.com/google/uuid"
	"os"
	"time"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		fmt.Printf("usage: %s <url>\n", cmd.Name)
		os.Exit(1)
	}
	url := cmd.Args[0]
	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return err
	}
	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return err
	}
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	feedFollowRow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}
	fmt.Println("Feed followed successfully")
	fmt.Printf("feed: %s\nuser: %s\n", feedFollowRow.FeedName, feedFollowRow.UserName)

	return nil

}

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		fmt.Println("Failed in getting current user!")
		return err
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		fmt.Println("Failed to get followed feeds for user")
		return err
	}
	for _, feed := range feedFollows {
		fmt.Printf("%s\n", feed.Name)
	}
	return nil

}
