package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"strconv"
)

func CommandBrowse(state *state, command command, user database.User) error {
	fmt.Println("Browse command received")

	limit := 2

	if len(command.Args) == 1 {

		arglimit, err := strconv.Atoi(command.Args[0])
		if err != nil {
			return fmt.Errorf("error parsing limit: %w", err)
		}
		limit = arglimit
	}

	posts, err := state.queries.GetPostsByUserId(context.Background(), database.GetPostsByUserIdParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}

	for _, post := range posts {
		fmt.Println("--------------------------------")
		fmt.Println(post.Title)
		fmt.Println("--------------------------------")
	}

	return nil
}
