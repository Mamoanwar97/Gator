package main

import (
	"database/sql"
	"fmt"
	"gator/internal/config"
	"os"

	"gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	configData, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
		return
	}
	db, err := sql.Open("postgres", configData.DbUrl)
	if err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
		return
	}
	defer db.Close()
	queries := database.New(db)
	newState := &state{config: &configData, queries: queries}
	commands := commands{commands: make(map[string]func(state *state, command command) error)}

	commands.register("login", CommandLogin)
	commands.register("register", CommandRegister)
	commands.register("reset", CommandReset)
	commands.register("users", CommandUsers)
	commands.register("agg", CommandAgg)
	commands.register("feeds", CommandFeeds)
	commands.register("addfeed", middlewareLoggedIn(CommandAddFeed))
	commands.register("follow", middlewareLoggedIn(CommandFollow))
	commands.register("following", middlewareLoggedIn(CommandFollowing))
	commands.register("unfollow", middlewareLoggedIn(CommandUnFollow))
	commands.register("browse", middlewareLoggedIn(CommandBrowse))

	cmdArg := os.Args

	if len(cmdArg) < 2 {
		fmt.Println("Usage: gator <command> <args...>")
		os.Exit(1)
		return
	}

	err = commands.run(newState, command{Command: cmdArg[1], Args: cmdArg[2:]})

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
		return
	}

	os.Exit(0)
}
