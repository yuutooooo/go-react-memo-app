package repository

import "github.com/yourusername/go-react-memo-app/internal/domain/model"

type NoteRepository interface {
	CreateNote(note *model.Note) (*model.Note, error)
	GetNoteByID(id string) (*model.Note, error)
	UpdateNote(note *model.Note) (*model.Note, error)
	DeleteNote(id string) error
	GetNotesByFolderID(userID string) ([]*model.Note, error)
	GetNotesByUserID(userID string) ([]*model.Note, error)
}
