package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yourusername/go-react-memo-app/internal/interface/dto"
	"github.com/yourusername/go-react-memo-app/internal/usecase"
)

type FolderController struct {
	folderUsecase usecase.FolderUsecase
}

func NewFolderController(folderUsecase usecase.FolderUsecase) *FolderController {
	return &FolderController{
		folderUsecase: folderUsecase,
	}
}

// フォルダ新規作成
func (c *FolderController) CreateFolder(ctx echo.Context) error {
	newFolder, err := dto.NewFolderRequest(ctx)
	log.Println("バリデーション後")
	log.Println(newFolder.ParentFolderID) //ok
	if err != nil {
		response := dto.ErrorResponse("リクエストの形式が正しくありません", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	userID := ctx.Get("user_id").(string)
	folderList, err := c.folderUsecase.CreateFolder(newFolder, userID)
	if err != nil {
		response := dto.FailResponse("フォルダの作成に失敗しました", err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	folderResponse := dto.CreateFolderResponseMany(folderList)
	response := dto.SuccessResponse(folderResponse, "フォルダを作成しました")
	return ctx.JSON(http.StatusCreated, response)
}

func (c *FolderController) GetFolderByID(ctx echo.Context) error {
	id := ctx.Param("id")
	updatedFolder, err := c.folderUsecase.GetFolderByID(id)
	if err != nil {
		response := dto.FailResponse("フォルダの取得に失敗しました", err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	folderResponse := dto.GetFolderByIDResponse(updatedFolder)
	response := dto.SuccessResponse(folderResponse, "フォルダを取得しました")
	return ctx.JSON(http.StatusOK, response)
}

func (c *FolderController) UpdateFolder(ctx echo.Context) error {
	id := ctx.Param("id")
	updatedFolder, err := dto.NewFolderRequest(ctx)
	if err != nil {
		response := dto.ErrorResponse("リクエストの形式が正しくありません", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	userID := ctx.Get("user_id").(string)
	updatedFolderList, err := c.folderUsecase.UpdateFolder(id, updatedFolder, userID)
	if err != nil {
		response := dto.FailResponse("フォルダの更新に失敗しました", err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	folderResponse := dto.CreateFolderResponseMany(updatedFolderList)
	response := dto.SuccessResponse(folderResponse, "フォルダを更新しました")
	return ctx.JSON(http.StatusOK, response)
}

func (c *FolderController) DeleteFolder(ctx echo.Context) error {
	id := ctx.Param("id")
	userID := ctx.Get("user_id").(string)
	deletedFolderList, err := c.folderUsecase.DeleteFolder(id, userID)
	if err != nil {
		response := dto.FailResponse("フォルダの削除に失敗しました", err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	folderResponse := dto.CreateFolderResponseMany(deletedFolderList)
	response := dto.SuccessResponse(folderResponse, "フォルダを削除しました")
	return ctx.JSON(http.StatusOK, response)
}
