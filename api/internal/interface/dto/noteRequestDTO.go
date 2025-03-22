package dto

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type NoteRequest struct {
	Title string `json:"title"`
	Content string `json:"content"`
	FolderID string `json:"folder_id"`
}

func NewNoteRequest(ctx echo.Context) (*NoteRequest, error) {
	var noteRequest NoteRequest
	if err := ctx.Bind(&noteRequest); err != nil {
		return nil, err
	}
	if err := noteRequest.Validate(); err != nil {
		return nil, err
	}
	return &noteRequest, nil
}

func (n *NoteRequest) Validate() error {
	if n.Title == "" {
		return errors.New("title is required")
	}
	if n.Content == "" {
		return errors.New("content is required")
	}
	if len(n.Content) > 100000 {
		return errors.New("content is too long")
	}
	if n.FolderID == "" {
		return errors.New("folder_id is required")
	}
	return nil
}
