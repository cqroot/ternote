package config

import (
	"fmt"
	"path"
	"strings"
	"sync"

	"github.com/adrg/xdg"
	"github.com/cqroot/ternote/pkg/types"
)

var (
	once sync.Once
)

func BasePath() (string, error) {
	once.Do(func() {
		xdg.DataFile("ternote/ternote.db")
	})

	dataDir, err := xdg.DataFile("ternote")
	if err != nil {
		return "", err
	}

	return dataDir, nil
}

func NotePath(id string) (string, error) {
	baseDir, err := BasePath()
	if err != nil {
		return "", err
	}
	return path.Join(baseDir, fmt.Sprintf("notes/%s.md", id)), nil
}

func VirtualPath(category string) (string, error) {
	basePath, err := BasePath()
	if err != nil {
		return "", err
	}
	return path.Join(basePath, "virtual", category), nil
}

func VirtualNotePath(note types.Note) (string, error) {
	categoryPath, err := VirtualPath(note.Category)
	if err != nil {
		return "", err
	}

	fileName := note.Title
	if strings.Trim(fileName, " ") == "" {
		fileName = "Untitled Note"
	}
	fileName = strings.Replace(fileName, "/", "|", -1)

	return path.Join(categoryPath, fmt.Sprintf("%s.md", fileName)),nil
}
