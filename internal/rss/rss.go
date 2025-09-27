package rss

import (
	"context"
	"fmt"
)

func RSS(feedURL string) (*RSSFeed, error) {
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching feed: %w", err)
	}
	return feed, nil
}
