package service

import (
	"github.com/yourusername/go-react-memo-app/internal/domain/model"
	"github.com/yourusername/go-react-memo-app/internal/domain/repository"
)

type NoteService interface {
	CreateNote(note *model.Note) (*model.Note, error)
	GetNoteByID(id string) (*model.Note, error)
	UpdateNote(note *model.Note) (*model.Note, error)
	DeleteNote(id string) error
}

type noteService struct {
	noteRepository repository.NoteRepository
}

func NewNoteService(noteRepository repository.NoteRepository) NoteService {
	return &noteService{
		noteRepository: noteRepository,
	}
}

func (s *noteService) CreateNote(note *model.Note) (*model.Note, error) {
	createNote, err := s.noteRepository.CreateNote(note)
	if err != nil {
		return nil, err
	}
	return createNote, nil
}

func (s *noteService) GetNoteByID(id string) (*model.Note, error) {
	note, err := s.noteRepository.GetNoteByID(id)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (s *noteService) UpdateNote(note *model.Note) (*model.Note, error) {
	updatedNote, err := s.noteRepository.UpdateNote(note)
	if err != nil {
		return nil, err
	}
	return updatedNote, nil
}

func (s *noteService) DeleteNote(id string) error {
	err := s.noteRepository.DeleteNote(id)
	if err != nil {
		return err
	}
	return nil
}
