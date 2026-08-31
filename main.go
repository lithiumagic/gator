package main

import (
	"fmt"
	"os"

	"github.com/lithiumagic/gator/internal/config"
)

// Update the main function to:

// Read the config file.
// Set the current user to "lane" (actually, you should use your name instead) and update the config file on disk.
// Read the config file again and print the contents of the config struct to the terminal.

func main() {
	cfg, err := config.Read()
	if err != nil {
		exitWithError(err)
	}
	err = cfg.SetUser("bob")
	if err != nil {
		exitWithError(err)
	}

	cfg, err = config.Read()
	if err != nil {
		exitWithError(err)
	}
	fmt.Printf("DbUrl: %s\n", cfg.DbUrl)
	fmt.Printf("CurrentUserName: %s\n", cfg.CurrentUserName)

}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
