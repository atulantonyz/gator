package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/atulantonyz/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	for _, item := range rssfeed.Channel.Item {

		layout := "Mon, 02 Jan 2006 15:04:05 -0700"
		parsedTime, err := time.Parse(layout, item.PubDate)
		if err != nil {
			fmt.Printf("Error parsing time %s: %v", item.PubDate, err)
			os.Exit(1)
			return fmt.Errorf("Error parsing date: %w", err)
		}

		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: len(item.Description) > 0},
			PublishedAt: parsedTime,
			FeedID:      feed.ID,
		}
		_, err = s.db.CreatePost(context.Background(), postParams)
		if err != nil {
			// Check if it's a unique violation
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" { // Unique violation
					continue
				}
			}
			fmt.Println("Some other error:", err)
			os.Exit(1)
			return fmt.Errorf("failed to create post: %w", err)
		}

	}
	return nil
}
