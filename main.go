package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/lithiumagic/gator/internal/config"
	"github.com/lithiumagic/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
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

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		exitWithError(err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	currentState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	commandsMap := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	commandsMap.register("login", handlerLogin)
	commandsMap.register("register", handlerRegisterUser)
	commandsMap.register("reset", handlerReset)
	commandsMap.register("users", handlerUsers)
	commandsMap.register("agg", handlerAgg)

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
