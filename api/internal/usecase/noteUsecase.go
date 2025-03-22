package usecase

import (
	"errors"

	"github.com/yourusername/go-react-memo-app/internal/domain/model"
	"github.com/yourusername/go-react-memo-app/internal/domain/service"
	"github.com/yourusername/go-react-memo-app/internal/interface/dto"
)

type NoteUsecase struct {
	noteService service.NoteService
}

func NewNoteUsecase(noteService service.NoteService) NoteUsecase {
	return NoteUsecase{
		noteService: noteService,
	}
}

func (u *NoteUsecase) CreateNote(note *dto.NoteRequest, userID string) (*model.Note, error) {
	noteModel := model.NewNote(note.Title, note.Content, note.FolderID, userID)
	createdNote, err := u.noteService.CreateNote(noteModel)
	if err != nil {
		return nil, err
	}
	return createdNote, nil
}

func (u *NoteUsecase) GetNoteByID(id string) (*model.Note, error) {
	note, err := u.noteService.GetNoteByID(id)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (u *NoteUsecase) UpdateNote(id string, userID string, note *dto.NoteRequest) (*model.Note, error) {
	registeredNote, err := u.noteService.GetNoteByID(id)
	if err != nil {
		return nil, err
	}
	if registeredNote.UserID() != userID {
		return nil, errors.New("unauthorized")
	}
	noteModel := model.NewUpdateNote(id, note.Title, note.Content, note.FolderID, userID, registeredNote.CreatedAt())
	updatedNote, err := u.noteService.UpdateNote(noteModel)
	if err != nil {
		return nil, err
	}
	return updatedNote, nil
}

func (u *NoteUsecase) DeleteNote(id string) error {
	err := u.noteService.DeleteNote(id)
	if err != nil {
		return err
	}
	return nil
}