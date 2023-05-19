package noterepo

import (
	"context"
	"sync"
	"time"

	"github.com/samarec1812/Snakerspace/internal/app"
	"github.com/samarec1812/Snakerspace/internal/notes"
)

type storage struct {
	data map[int64]*notes.Note
	mu   *sync.RWMutex
	inc  int64
}

func New() app.NoteRepository {
	return &storage{
		data: make(map[int64]*notes.Note),
		mu:   &sync.RWMutex{},
		inc:  -1,
	}
}

func (s *storage) Create(_ context.Context, note notes.Note) (notes.Note, error) {

	s.mu.Lock()
	s.inc++
	newNote := &notes.Note{
		ID:        s.inc,
		Text:      note.Text,
		CreatedAt: time.Now().UTC(),
		UpdateAt:  time.Now().UTC(),
	}

	s.data[s.inc] = newNote
	s.mu.Unlock()

	return *newNote, nil
}

func (s *storage) GetByID(_ context.Context, noteID int64) (notes.Note, error) {
	s.mu.RLock()
	val, ok := s.data[noteID]
	s.mu.RUnlock()

	if !ok {
		return notes.Note{}, notes.ErrNoteNotFound
	}

	return *val, nil
}

func (s *storage) Update(_ context.Context, noteID int64, noteText string) (notes.Note, error) {
	s.mu.RLock()
	val, ok := s.data[noteID]
	s.mu.RUnlock()

	if !ok {
		return notes.Note{}, notes.ErrNoteNotFound
	}

	val.Text = noteText
	val.UpdateAt = time.Now().UTC()

	s.mu.Lock()
	s.data[noteID] = val
	s.mu.Unlock()

	return *val, nil
}
func (s *storage) Delete(_ context.Context, noteID int64) error {
	s.mu.RLock()
	_, ok := s.data[noteID]
	s.mu.RUnlock()

	if !ok {
		return notes.ErrNoteNotFound
	}

	s.mu.Lock()
	delete(s.data, noteID)
	s.mu.Unlock()

	return nil
}

func (s *storage) GetAll(_ context.Context) ([]notes.Note, error) {
	var list []notes.Note
	for _, note := range s.data {
		s.mu.RLock()
		list = append(list, *note)
		s.mu.RUnlock()
	}

	return list, nil
}
