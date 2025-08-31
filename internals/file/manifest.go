package file

import (
	"encoding/json"
	"errors"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type File struct {
	filepath string

	Name         string
	Size         int64
	MimeType     string
	LastModified time.Time
}

func fileFromPath(path string) (File, error) {
	file := File{}

	stat, err := os.Stat(path)
	if err != nil {
		return file, err
	}

	file.filepath = path
	file.Name = stat.Name()
	file.Size = stat.Size()
	file.MimeType = mime.TypeByExtension(filepath.Ext(path))
	file.LastModified = stat.ModTime().UTC()
	return file, nil
}

func (f File) Open() (io.ReadCloser, error) {
	return os.Open(f.filepath)
}

type Manifest struct {
	entryPointMap map[string]string
	outfileMap    map[string]File
}

func (m *Manifest) GetUrls(paths ...string) []string {
	result := make([]string, 0, len(paths))

	for _, path := range paths {
		if entry, ok := m.entryPointMap[path]; ok {
			result = append(result, entry)
		}
	}

	return result
}

func (m *Manifest) GetFile(path string) (File, error) {
	file, ok := m.outfileMap[path]
	if !ok {
		return File{}, errors.New("not found")
	}

	return file, nil
}

func ReadManifest(reader io.Reader) (*Manifest, error) {
	type manifestEntry struct {
		Entrypoint string `json:"entrypoint"`
		Outfile    string `json:"outfile"`
	}

	type manifestList []manifestEntry

	list := make(manifestList, 0)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&list)
	if err != nil {
		return nil, err
	}

	manifest := new(Manifest)
	manifest.entryPointMap = make(map[string]string, len(list))
	manifest.outfileMap = make(map[string]File, len(list))

	for _, entry := range list {
		entrypoint := strings.TrimPrefix(entry.Entrypoint, "assets/")
		outfile := strings.TrimPrefix(entry.Outfile, "dist/")

		file, err := fileFromPath(outfile)
		if err != nil {
			return nil, err
		}
		manifest.entryPointMap[entrypoint] = outfile
		manifest.outfileMap[outfile] = file
	}

	return manifest, nil
}
