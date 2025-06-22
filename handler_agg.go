package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		fmt.Printf("usage: %s <url>\n", cmd.Name)
		os.Exit(1)
	}
	time_between_reqs := cmd.Args[0]
	timeBetweenReqs, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("failed to parse time between requests: %w", err)
	}
	fmt.Printf("Collecting feeds every %s\n", time_between_reqs)
	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
		fmt.Println("Feed fetched successfully at: ", time.Now())
	}
	return nil
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("could not get next feed to fetch: %w", err)
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("could not mark fetched feed: %w", err)
	}
	rssfeed, err := fetchFeed(context.Background(), feed.Url)
	for _, item := range rssfeed.Channel.Item {
		fmt.Printf("%s\n", item.Title)
	}
	return nil
}
