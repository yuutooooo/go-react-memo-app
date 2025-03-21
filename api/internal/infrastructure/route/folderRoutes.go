package route

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/yourusername/go-react-memo-app/internal/infrastructure/config"
	"github.com/yourusername/go-react-memo-app/internal/infrastructure/di"
	"github.com/yourusername/go-react-memo-app/internal/infrastructure/middleware"
	"github.com/yourusername/go-react-memo-app/internal/interface/controller"
)

func NewFolderController() *controller.FolderController {
	db, err := config.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	container := di.NewContainer(db)
	return container.FolderController()
}

func SetupFolderRoutes(e *echo.Group) {
	folderController := NewFolderController()
	auth := e.Group("")
	auth.Use(middleware.AuthMiddleware)
	auth.POST("", folderController.CreateFolder)
	auth.GET("/:id", folderController.GetFolderByID)
	auth.PUT("/:id", folderController.UpdateFolder)
	auth.DELETE("/:id", folderController.DeleteFolder)
}