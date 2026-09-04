package main

import (
	"fmt"
	"os"

	"github.com/lithiumagic/gator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name      string
	arguments []string
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		exitWithError(err)
	}

	currentState := &state{
		config: &cfg,
	}

	commandsMap := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	commandsMap.register("login", handlerLogin)

	if len(os.Args) < 2 {
		err = fmt.Errorf("usage: gator <command> [args...]")
		exitWithError(err)
	}

	comm := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}
	err = commandsMap.run(currentState, comm)
	if err != nil {
		exitWithError(err)
	}
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
