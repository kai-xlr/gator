package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {

}

func Read() (Config, error) {

}

func getConfigFilePath() (string, error) {

}

func write(cfg Config) error {

}
