package assets

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

//go:embed dist/assets/*
var assets embed.FS

var manifest Manifest

var (
	NotFound             = errors.New("asset not found")
	EncodingDoesNotExist = errors.New("encoding does not exist")
)

func init() {
	type manifestEntryVersion struct {
		Filepath string `json:"filepath"`
		Encoding string `json:"encoding"`
		Size     int64  `json:"size"`
	}

	type manifestEntry struct {
		Entrypoint string                 `json:"entrypoint"`
		Mimetype   string                 `json:"mimetype"`
		Versions   []manifestEntryVersion `json:"versions"`
	}

	type manifestEntrySlice []manifestEntry

	entries := manifestEntrySlice{}

	manifestFile, err := os.Open("manifest.json")
	if err != nil {
		panic(fmt.Errorf("error loading manifest: %w", err))
	}
	defer manifestFile.Close()

	err = json.NewDecoder(manifestFile).Decode(&entries)
	if err != nil {
		panic(fmt.Errorf("error loading manifest: %w", err))
	}

	manifest = Manifest{}
	manifest.entryPointMap = make(map[string]string, len(entries))
	manifest.outfileMap = make(map[string]Asset, len(entries))

	for _, e := range entries {
		asset := Asset{}
		asset.Name = filepath.Base(e.Entrypoint)
		asset.Mimetype = e.Mimetype
		asset.preencoded = make(map[string]File, len(e.Versions)-1)

		for _, version := range e.Versions {
			file := File{}
			file.filepath = version.Filepath
			file.Size = version.Size
			file.Encoding = version.Encoding

			if version.Encoding == "none" {
				asset.fallback = file
			} else {
				asset.preencoded[version.Encoding] = file
			}
		}

		manifest.entryPointMap[e.Entrypoint] = asset.fallback.filepath
		manifest.outfileMap[asset.fallback.filepath] = asset
	}
}

func GetManifest() *Manifest {
	return &manifest
}

type Asset struct {
	Name     string
	Mimetype string

	fallback   File
	preencoded map[string]File
}

func (a Asset) GetBase() File {
	return a.fallback
}

func (a Asset) GetEncodedFile(encoding string) (File, error) {
	file, ok := a.preencoded[encoding]
	if !ok {
		return File{}, EncodingDoesNotExist
	}

	return file, nil
}

func (a Asset) GetBestCandidate() File {
	bestcandidate := a.fallback

	for _, encoded := range a.preencoded {
		if encoded.Size < bestcandidate.Size {
			bestcandidate = encoded
		}
	}

	return bestcandidate
}

type File struct {
	filepath string
	Size     int64
	Encoding string
}

func (f File) Open() (io.ReadCloser, error) {
	return assets.Open("dist/" + f.filepath)
}

type Manifest struct {
	entryPointMap map[string]string
	outfileMap    map[string]Asset
}

func (m *Manifest) MapEntrypoints(entrypoints ...string) []string {
	result := make([]string, 0, len(entrypoints))

	for _, path := range entrypoints {
		if entry, ok := m.entryPointMap[path]; ok {
			result = append(result, entry)
		}
	}

	return result
}

func (m *Manifest) GetAsset(path string) (Asset, error) {
	fmt.Println(path)
	fmt.Println(m.outfileMap)
	file, ok := m.outfileMap[path]
	if !ok {
		return Asset{}, NotFound
	}

	return file, nil
}
