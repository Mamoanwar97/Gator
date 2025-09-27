package main

import "fmt"

type commands struct {
	commands map[string]func(state *state, command command) error
}

func (c *commands) run(s *state, cmd command) error {
	command, ok := c.commands[cmd.Command]
	if !ok {
		return fmt.Errorf("command not found: %s", cmd.Command)
	}
	return command(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}
