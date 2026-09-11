package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

// Define a struct
type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Define a method on Config
func (cfg *Config) SetUser(userName string) error {
	// Export a SetUser method on the Config struct that writes the config struct to the JSON file after setting the current_user_name field.
	cfg.CurrentUserName = userName
	return write(cfg)

}

func Read() (Config, error) {
	// Export a Read function that reads the JSON file found at ~/.gatorconfig.json and returns a Config struct.
	// It should read the file from the HOME directory, then decode the JSON string into a new Config struct.
	// I used os.UserHomeDir to get the location of HOME.
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(path) // Read the file into []byte with os.ReadFile.
	if err != nil {
		return Config{}, err
	}

	var cfg Config // create empty Config
	// err = json.Unmarshal(data, &cfg) // Decode the bytes into it with json.Unmarshal
	// if err != nil {
	// 	return Config{}, err
	// }
	if err := json.Unmarshal(data, &cfg); err != nil { // same as above, but idiomatic Go
		return Config{}, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, configFileName), nil
}

func write(cfg *Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
