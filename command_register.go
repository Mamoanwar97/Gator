package main

import (
	"context"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func CommandRegister(state *state, command command) error {
	fmt.Println("Register command received")

	if len(command.Args) != 1 {
		return fmt.Errorf("register command requires exactly 1 argument")
	}

	user := command.Args[0]
	_, err := state.queries.CreateUser(context.Background(), database.CreateUserParams{
		Name: user,
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	
	config.SetUser(state.config, user)

	fmt.Println("User created:", user)

	return nil
}
