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
		response := dto.ErrorResponse("リクエストの形式が正しくありません", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	userID := ctx.Get("user_id").(string)
	note, err := c.noteUsecase.CreateNote(newNote, userID)
	if err != nil {
		response := dto.FailResponse("メモの作成に失敗しました", err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	noteResponse := dto.CreateNoteResponse(note)
	response := dto.SuccessResponse(noteResponse, "メモを作成しました")
	return ctx.JSON(http.StatusCreated, response)
}

func (c *NoteController) GetNoteByID(ctx echo.Context) error {
	id := ctx.Param("id")
	note, err := c.noteUsecase.GetNoteByID(id)
	if err != nil {
		response := dto.FailResponse("メモの取得に失敗しました", err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	noteResponse := dto.CreateNoteResponse(note)
	response := dto.SuccessResponse(noteResponse, "メモを取得しました")
	return ctx.JSON(http.StatusOK, response)
}

func (c *NoteController) UpdateNote(ctx echo.Context) error {
	id := ctx.Param("id")
	userID := ctx.Get("user_id").(string)
	updateNote, err := dto.NewNoteRequest(ctx)
	if err != nil {
		response := dto.ErrorResponse("リクエストの形式が正しくありません", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	note, err := c.noteUsecase.UpdateNote(id, userID, updateNote)
	if err != nil {
		response := dto.FailResponse("メモの更新に失敗しました", err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	noteResponse := dto.CreateNoteResponse(note)
	response := dto.SuccessResponse(noteResponse, "メモを更新しました")
	return ctx.JSON(http.StatusOK, response)
}

func (c *NoteController) DeleteNote(ctx echo.Context) error {
	id := ctx.Param("id")
	err := c.noteUsecase.DeleteNote(id)
	if err != nil {
		response := dto.FailResponse("メモの削除に失敗しました", err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response := dto.SuccessResponse(nil, "メモを削除しました")
	return ctx.JSON(http.StatusOK, response)
}

func (c *NoteController) GetNotesByUserID(ctx echo.Context) error {
	userID := ctx.Get("user_id").(string)
	notes, err := c.noteUsecase.GetNotesByUserID(userID)
	if err != nil {
		response := dto.FailResponse("メモ一覧の取得に失敗しました", err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	notesResponse := dto.CreateNoteResponseMany(notes)
	response := dto.SuccessResponse(notesResponse, "メモ一覧を取得しました")
	return ctx.JSON(http.StatusOK, response)
}
