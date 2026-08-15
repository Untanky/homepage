package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

func Load(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("loading config: %w", err)
	}

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, fmt.Errorf("decoding config: %w", err)
	}

	return config, nil
}
