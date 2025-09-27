package main

import (
	"context"
	"fmt"
)

func CommandUsers(state *state, command command) error {
	data, err := state.queries.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %w", err)
	}

	currentUser := state.config.CurrentUserName

	for _, user := range data {
		if user.Name == currentUser {
			fmt.Println("*", user.Name, "(current)")
		} else {
			fmt.Println("*", user.Name)
		}
	}

	return nil
}
