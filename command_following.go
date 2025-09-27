package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func CommandFollowing(state *state, command command, user database.User) error {
	feedFollows, err := state.queries.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return fmt.Errorf("error getting feed: %w", err)
	}

	for _, feedFollow := range feedFollows {
		fmt.Println(feedFollow.FeedName)
	}

	fmt.Println()

	return nil
}
