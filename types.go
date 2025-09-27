package main

import (
	"gator/internal/config"
	"gator/internal/database"
)

type state struct {
	config *config.Config
	queries *database.Queries
}

type command struct {
	Command string
	Args []string
}