package ternote

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/cqroot/ternote/internal/metadata"
	"github.com/cqroot/ternote/pkg/config"
	"github.com/cqroot/ternote/pkg/types"
)

type Ternote struct {
}

func New() *Ternote {
	return &Ternote{}
}

func (t Ternote) Notes() []types.Note {
	return metadata.Notes()
}

func (t Ternote) UpdateNoteMetadata(note *types.Note) error {
	notePath, err := config.NotePath(note.Id)
	if err != nil {
		return err
	}

	fi, err := os.Stat(notePath)
	if err != nil {
		return fmt.Errorf("%w: %s", types.ErrorNoteFileNoteFound, notePath)
	}
	modTime := fi.ModTime()
	if modTime.Equal(note.ModTime) {
		return nil
	}

	var newTitle string = ""

	if fi.Size() != 0 {
		file, err := os.Open(notePath)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Scan()
		firstLine := scanner.Text()
		if err := scanner.Err(); err != nil {
			return err
		}

		if strings.HasPrefix(firstLine, "# ") && len(firstLine) > 2 {
			newTitle = firstLine[2:]
		}
	}

	note.Title = newTitle
	note.ModTime = modTime
	metadata.Update(note)
	return nil
}

func (t Ternote) NewNote(category string) error {
	id := time.Now().Format("20060102150405")
	notePath, err := config.NotePath(id)
	if err != nil {
		return err
	}

	// Create new file
	f, err := os.Create(notePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Update metadata
	fi, err := f.Stat()
	if err != nil {
		return err
	}

	note := types.Note{
		Id:       id,
		Category: category,
		Title:    "",
		ModTime:  fi.ModTime(),
	}

	metadata.Update(&note)

	return nil
}

func (t Ternote) RemoveNote(id string) error {
	basePath, err := config.BasePath()
	if err != nil {
		return nil
	}
	notePath, err := config.NotePath(id)
	if err != nil {
		return nil
	}
	trashPath := path.Join(basePath, "trash")

	if _, err := os.Stat(trashPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(trashPath, os.ModePerm)
		if err != nil {
			return nil
		}
	}

	err = os.Rename(notePath, path.Join(trashPath, path.Base(notePath)))
	if err != nil {
		return err
	}

	metadata.RemoveNote(id)

	return nil
}
