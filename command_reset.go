package main

import (
	"context"
	"fmt"
)

func CommandReset(state *state, command command) error {
	err := state.queries.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting users: %w", err)
	}
	err = state.queries.DeleteFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error deleting feeds: %w", err)
	}
	return nil
}
