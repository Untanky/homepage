package config

import "github.com/untanky/homepage/internal/database"

type Config struct {
	Database database.Config `yaml:"database"`
}
