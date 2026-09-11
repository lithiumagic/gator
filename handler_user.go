package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lithiumagic/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("No arguments were given. Usage: login <username>")
	}

	_, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Printf("%s has been set\n", cmd.arguments[0])
	return nil

}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.arguments) > 0 {
		return errors.New("Too many arguments were given. Usage: users")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Couldn't get users. Error: %w", err)
	}

	for _, user := range users {
		if user.Name != s.cfg.CurrentUserName {
			fmt.Printf("* %s\n", user.Name)
		} else {
			fmt.Printf("* %s (current)\n", user.Name)
		}

	}

	return nil

}

func handlerRegisterUser(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("No arguments were given. Usage: login <username>")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
	})
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("the user: %v was created", user)
	log.Printf("the user: %v", user)
	return nil

}

func handlerReset(s *state, cmd command) error {
	if len(cmd.arguments) > 0 {
		return errors.New("Too many arguments were given. Usage: reset")
	}

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset the database: %w", err)
	}

	fmt.Println("all users from users table were deleted")
	return nil

}
