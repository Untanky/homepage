package config

import "go.lukasgrimm.me/homepage/internal/database"

type Config struct {
	Database database.Config `yaml:"database"`
}
