package app

import (
	"context"
	"fmt"
	"sort"

	"github.com/samarec1812/Snakerspace/internal/notes"
)

type App interface {
	CreateNote(context.Context, string) (notes.Note, error)
	GetNote(context.Context, int64) (notes.Note, error)
	UpdateNote(context.Context, int64, string) (notes.Note, error)
	DeleteNote(context.Context, int64) error
	GetAllNote(context.Context, string) ([]notes.Note, error)
}

type app struct {
	noterepo NoteRepository
}

func (a *app) CreateNote(ctx context.Context, text string) (notes.Note, error) {
	note := notes.Note{
		Text: text,
	}
	if err := note.Validate(); err != nil {
		return notes.Note{}, fmt.Errorf("error with validation: %w", err)
	}
	newNote, err := a.noterepo.Create(ctx, note)
	if err != nil {
		return notes.Note{}, fmt.Errorf("error with create: %w", err)
	}

	return newNote, nil
}

func (a *app) GetNote(ctx context.Context, noteID int64) (notes.Note, error) {
	note, err := a.noterepo.GetByID(ctx, noteID)
	if err != nil {
		return notes.Note{}, fmt.Errorf("error with get not by ID: %w", err)
	}

	return note, nil
}

func (a *app) UpdateNote(ctx context.Context, noteID int64, noteText string) (notes.Note, error) {
	if noteText == "" {
		return notes.Note{}, notes.ErrInvalidParam
	}

	note, err := a.noterepo.Update(ctx, noteID, noteText)
	if err != nil {
		return notes.Note{}, fmt.Errorf("error with update note: %w", err)
	}

	return note, nil
}

func (a *app) DeleteNote(ctx context.Context, noteID int64) error {
	err := a.noterepo.Delete(ctx, noteID)
	if err != nil {
		return fmt.Errorf("error with delete note: %w", err)
	}

	return nil
}

func (a *app) GetAllNote(ctx context.Context, param string) ([]notes.Note, error) {
	list, err := a.noterepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error with get list note: %w", err)
	}

	switch param {
	case "id":
		sort.Slice(list, func(i, j int) bool {
			return list[i].ID < list[j].ID
		})
	case "text":
		sort.Slice(list, func(i, j int) bool {
			return list[i].Text < list[j].Text
		})
	case "created_at":
		sort.Slice(list, func(i, j int) bool {
			return list[i].CreatedAt.Unix() < list[j].CreatedAt.Unix()
		})
	case "updated_at":
		sort.Slice(list, func(i, j int) bool {
			return list[i].UpdateAt.Unix() < list[j].UpdateAt.Unix()
		})
	}

	return list, nil
}

type NoteRepository interface {
	Create(context.Context, notes.Note) (notes.Note, error)
	GetByID(context.Context, int64) (notes.Note, error)
	Update(context.Context, int64, string) (notes.Note, error)
	Delete(context.Context, int64) error
	GetAll(context.Context) ([]notes.Note, error)
}

func NewApp(noteRepo NoteRepository) App {
	return &app{
		noterepo: noteRepo,
	}
}
