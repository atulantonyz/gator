package main

import (
	"context"
	"fmt"
	"os"
)

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	rssfeed, err := fetchFeed(context.Background(), url)
	if err != nil {
		fmt.Println("Unable to fetch feed")
		os.Exit(1)
	}
	fmt.Println(*rssfeed)
	return nil
}
