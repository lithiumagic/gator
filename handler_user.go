package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("No arguments were given. Usage: login <username>")
	}

	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Printf("%s has been set\n", cmd.arguments[0])
	return nil

}
