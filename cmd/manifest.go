package main

import (
	"encoding/json"
	"os"
	"slices"
	"strings"

	"github.com/untanky/homepage/internals/handlers"
)

type manifestEntry struct {
	Entrypoint string `json:entrypoint`
	Outfile    string `json:outfile`
}

type manifest []manifestEntry

func (m manifest) GetUrl(path string) string {
	index := slices.IndexFunc(m, func(entry manifestEntry) bool {
		return strings.EqualFold(path, entry.Entrypoint)
	})

	if index == -1 {
		return ""
	}

	return m[index].Outfile
}

func readManifest() (handlers.Manifest, error) {
	file, err := os.Open("manifest.json")
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	manifest := make(manifest, 0)
	err = decoder.Decode(&manifest)
	if err != nil {
		return nil, err
	}

	for i, v := range manifest {
		manifest[i] = manifestEntry{
			Entrypoint: strings.TrimPrefix(v.Entrypoint, "assets/"),
			Outfile:    strings.TrimPrefix(v.Outfile, "dist"),
		}
	}

	return manifest, nil
}
