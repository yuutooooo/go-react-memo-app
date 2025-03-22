package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yourusername/go-react-memo-app/internal/interface/dto"
	"github.com/yourusername/go-react-memo-app/internal/usecase"
)

type NoteController struct {
	noteUsecase usecase.NoteUsecase
}

func NewNoteController(noteUsecase usecase.NoteUsecase) *NoteController {
	return &NoteController{
		noteUsecase: noteUsecase,
	}
}

func (c *NoteController) CreateNote(ctx echo.Context) error {
	newNote, err := dto.NewNoteRequest(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	userID := ctx.Get("user_id").(string)
	note, err := c.noteUsecase.CreateNote(newNote, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := dto.CreateNoteResponse(note)
	return ctx.JSON(http.StatusCreated, response)
}

func (c *NoteController) GetNoteByID(ctx echo.Context) error {
	id := ctx.Param("id")
	note, err := c.noteUsecase.GetNoteByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := dto.CreateNoteResponse(note)
	return ctx.JSON(http.StatusOK, response)
}

func (c *NoteController) UpdateNote(ctx echo.Context) error {
	id := ctx.Param("id")
	userID := ctx.Get("user_id").(string)
	updateNote, err := dto.NewNoteRequest(ctx)
	note, err := c.noteUsecase.UpdateNote(id, userID, updateNote)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := dto.CreateNoteResponse(note)
	return ctx.JSON(http.StatusOK, response)
}

func (c *NoteController) DeleteNote(ctx echo.Context) error {
	id := ctx.Param("id")
	err := c.noteUsecase.DeleteNote(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, "Note deleted successfully")
}
