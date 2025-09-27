package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func CommandUnFollow(state *state, command command, user database.User) error {
	fmt.Println("Follow command received")

	if len(command.Args) != 1 {
		return fmt.Errorf("follow command requires exactly 1 argument")
	}

	url := command.Args[0]

	feed, err := state.queries.GetFeedByURL(context.Background(), url)

	if err != nil {
		return fmt.Errorf("error getting feed: %w", err)
	}

	err = state.queries.DeleteFeedFollows(context.Background(), database.DeleteFeedFollowsParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("error deleting feed follow: %w", err)
	}

	return nil
}