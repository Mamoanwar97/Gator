package main

import (
	"context"
	"fmt"
	"gator/internal/config"
)

func CommandLogin(state *state, command command) error {
	fmt.Println("Login command received")

	if len(command.Args) != 1 {
		return fmt.Errorf("login command requires exactly 1 argument")
	}

	user := command.Args[0]
	_, err := state.queries.GetUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}
	err = config.SetUser(state.config, user)
	if err != nil {
		return fmt.Errorf("error setting user: %w", err)
	}

	fmt.Println("User set:", user)

	return nil
}
