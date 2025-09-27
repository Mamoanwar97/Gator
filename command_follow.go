package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func CommandFollow(state *state, command command, user database.User) error {
	fmt.Println("Follow command received")

	if len(command.Args) != 1 {
		return fmt.Errorf("follow command requires exactly 1 argument")
	}

	url := command.Args[0]

	feed, err := state.queries.GetFeedByURL(context.Background(), url)

	if err != nil {
		return fmt.Errorf("error getting feed: %w", err)
	}

	feedFollow, err := state.queries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Println("Feed follow created:", feedFollow.FeedName, feedFollow.UserName)

	return nil
}