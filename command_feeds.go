package main

import (
	"context"
	"fmt"
)

func CommandFeeds(state *state, command command) error {
	feeds, err := state.queries.GetFeeds(context.Background())

	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println()
	}

	return nil
}
