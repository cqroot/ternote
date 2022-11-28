package types

import (
	"errors"
	"fmt"
)

var (
	ErrorBase = errors.New("Ternote Error")

	ErrorNoteFileNoteFound = fmt.Errorf("%w: note file not found", ErrorBase)
)
