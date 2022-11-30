package virtual

import (
	"os"
	"path"

	"github.com/cqroot/ternote/pkg/config"
	"github.com/cqroot/ternote/pkg/types"
)

func NewNote(note types.Note) error {
	categoryPath, err := config.VirtualPath(note.Category)
	if err != nil {
		return err
	}

	err = os.MkdirAll(categoryPath, os.ModePerm)
	if err != nil {
		return err
	}

	notePath, err := config.NotePath(note.Id)
	if err != nil {
		return err
	}

	virtualNotePath, err := config.VirtualNotePath(note)
	if err != nil {
		return err
	}

	os.Symlink(
		notePath,
		virtualNotePath,
	)

	return nil
}

func RemoveNote(note types.Note) error {
	virtualNotePath, err := config.VirtualNotePath(note)
	if err != nil {
		return err
	}

	basePath, err := config.BasePath()
	if err != nil {
		return nil
	}
	trashPath := path.Join(basePath, "trash")

	return os.Rename(virtualNotePath, path.Join(trashPath, path.Base(virtualNotePath)))
}
