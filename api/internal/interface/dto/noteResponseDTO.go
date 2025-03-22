package dto

import (
	"time"

	"github.com/yourusername/go-react-memo-app/internal/domain/model"
)

type NoteResponse struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	FolderID string `json:"folder_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewNoteResponse(note *model.Note) *NoteResponse {
	return &NoteResponse{
		ID: note.ID(),
		Title: note.Title(),
		Content: note.Content(),
		FolderID: note.FolderID(),
		CreatedAt: note.CreatedAt(),
		UpdatedAt: note.UpdatedAt(),
	}
}

func CreateNoteResponse(note *model.Note) *NoteResponse {
	return NewNoteResponse(note)
}

func CreateNoteResponseMany(noteList []*model.Note) []*NoteResponse {
	noteResponses := make([]*NoteResponse, len(noteList))
	for i, note := range noteList {
		noteResponses[i] = NewNoteResponse(note)
	}
	return noteResponses
}
