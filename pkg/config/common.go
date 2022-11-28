package config

import (
	"fmt"
	"path"
	"sync"

	"github.com/adrg/xdg"
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
