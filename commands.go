package main 

import (
	"errors"
)

type command struct {
	Name string
	Args	[]string
}

type commands struct {
	regCommand map[string]func(*state, command)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.regCommand[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	exec, ok := c.regCommand[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return exec(s, cmd)
}


