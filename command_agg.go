package main

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"gator/internal/rss"
	"strings"
	"time"

	"github.com/google/uuid"
)

func scrapeFeeds(state *state) error {
	feeds, err := state.queries.GetNextFeedToFetch(context.Background())

	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %w", err)
	}

	err = state.queries.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            feeds.ID,
	})

	if err != nil {
		return fmt.Errorf("error marking feed fetched: %w", err)
	}

	feed, err := rss.RSS(feeds.Url)

	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	for _, item := range feed.Channel.Item {
		publishedAt, err := time.Parse(time.RFC1123, item.PubDate)

		if err != nil {
			publishedAt = time.Now()
		}

		_, err = state.queries.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			PublishedAt: publishedAt,
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			FeedID:      feeds.ID,
		})

		if err != nil {
			// Check if it's a duplicate URL error (unique constraint violation)
			if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "UNIQUE constraint") {
				// Ignore duplicate URL errors and continue processing
				fmt.Printf("Skipping duplicate post: %s\n", item.Title)
				continue
			}
			// Log other errors but continue processing
			fmt.Printf("Error creating post '%s': %v\n", item.Title, err)
			continue
		}

		fmt.Println("--------------------------------")
		fmt.Println(item.Title)
		fmt.Println("--------------------------------")
	}

	return nil
}

func CommandAgg(state *state, command command) error {
	fmt.Println("Agg command received")

	if len(command.Args) != 1 {
		return fmt.Errorf("agg command requires exactly 1 argument")
	}

	timeBetweenRequests, err := time.ParseDuration(command.Args[0])

	if err != nil {
		return fmt.Errorf("error parsing time between requests: %w", err)
	}

	fmt.Println("Collecting feeds every", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(state)
	}
}
