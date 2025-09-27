package main

import (
	"context"
	"gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(state *state, command command) error {
		user, err := state.queries.GetUser(context.Background(), state.config.CurrentUserName)

		if err != nil {
			return err
		}

		return handler(state, command, user)
	}
}
