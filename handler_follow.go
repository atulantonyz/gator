package main

import (
	"context"
	"fmt"
	"github.com/atulantonyz/gator/internal/database"
	"github.com/google/uuid"
	"os"
	"time"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		fmt.Printf("usage: %s <url>\n", cmd.Name)
		os.Exit(1)
	}
	url := cmd.Args[0]
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

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		fmt.Println("Failed to get followed feeds for user")
		return err
	}
	for _, feed := range feedFollows {
		fmt.Printf("%s\n", feed.Name)
	}
	return nil

}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		fmt.Printf("usage: %s <url>\n", cmd.Name)
		os.Exit(1)
	}
	url := cmd.Args[0]

	deleteFollowParams := database.DeleteFeedFollowForUserParams{
		UserID: user.ID,
		Url:    url,
	}
	err := s.db.DeleteFeedFollowForUser(context.Background(), deleteFollowParams)
	if err != nil {
		return err
	}
	return nil
}
