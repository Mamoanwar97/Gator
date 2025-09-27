package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func CommandAddFeed(state *state, command command, user database.User) error {
	fmt.Println("Add feed command received")

	if len(command.Args) != 2 {
		return fmt.Errorf("add feed command requires exactly 2 arguments")
	}

	name := command.Args[0]
	url := command.Args[1]

	feed, err := state.queries.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
	})

	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	_, err = state.queries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Println(feed)

	return nil
}
