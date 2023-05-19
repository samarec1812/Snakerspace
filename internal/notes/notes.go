package notes

import (
	"errors"
	"time"
)

var (
	ErrInvalidNote  = errors.New("invalid note")
	ErrNoteNotFound = errors.New("note not found")
	ErrInvalidParam = errors.New("invalid input param")
)

type Note struct {
	ID        int64
	Text      string
	CreatedAt time.Time
	UpdateAt  time.Time
}

func (n Note) Validate() error {
	if n.Text == "" {
		return ErrInvalidNote
	}

	return nil
}
