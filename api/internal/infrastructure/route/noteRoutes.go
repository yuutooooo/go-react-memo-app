package route

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/yourusername/go-react-memo-app/internal/infrastructure/config"
	"github.com/yourusername/go-react-memo-app/internal/infrastructure/di"
	"github.com/yourusername/go-react-memo-app/internal/infrastructure/middleware"
	"github.com/yourusername/go-react-memo-app/internal/interface/controller"
)

func NewNoteController() *controller.NoteController {
	db, err := config.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	container := di.NewContainer(db)
	return container.NoteController()
}

func SetupNoteRoutes(e *echo.Group) {
	noteController := NewNoteController()
	auth := e.Group("")
	auth.Use(middleware.AuthMiddleware)
	auth.POST("", noteController.CreateNote)
	auth.GET("/:id", noteController.GetNoteByID)
	auth.PUT("/:id", noteController.UpdateNote)
	auth.DELETE("/:id", noteController.DeleteNote)
}
