package main

import "fmt"

type commands struct {
	handlers map[string]func(*state, command) error
}

// This method runs a given command with the provided state if it exists.
func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %v", cmd.name)
	}

	return handler(s, cmd)

}

func (c *commands) register(name string, f func(*state, command) error) { // - This method registers a new handler function for a command name
	c.handlers[name] = f
}
