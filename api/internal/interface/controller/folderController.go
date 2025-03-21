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
	log.Println(newFolder.ParentFolderID)//ok
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	userID := ctx.Get("user_id").(string)
	folderList, err := c.folderUsecase.CreateFolder(newFolder, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := dto.CreateFolderResponseMany(folderList)
	return ctx.JSON(http.StatusCreated, response)
}

func (c *FolderController) GetFolderByID(ctx echo.Context) error {
	id := ctx.Param("id")
	updatedFolder, err := c.folderUsecase.GetFolderByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := dto.GetFolderByIDResponse(updatedFolder)
	return ctx.JSON(http.StatusOK, response)
}

func (c *FolderController) UpdateFolder(ctx echo.Context) error {
	id := ctx.Param("id")
	updatedFolder, err := dto.NewFolderRequest(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	userID := ctx.Get("user_id").(string)
	updatedFolderList, err := c.folderUsecase.UpdateFolder(id, updatedFolder, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := dto.CreateFolderResponseMany(updatedFolderList)
	return ctx.JSON(http.StatusOK, response)
}

func (c *FolderController) DeleteFolder(ctx echo.Context) error {
	id := ctx.Param("id")
	userID := ctx.Get("user_id").(string)
	deletedFolderList, err := c.folderUsecase.DeleteFolder(id, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	response := dto.CreateFolderResponseMany(deletedFolderList)
	return ctx.JSON(http.StatusOK, response)
}
