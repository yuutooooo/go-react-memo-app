package route

import (
	"log"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	log.Println("Setting up routes")
	userRouter := e.Group("/user")
	SetupUserRoutes(userRouter)
	folderRouter := e.Group("/folder")
	SetupFolderRoutes(folderRouter)
	noteRouter := e.Group("/note")
	SetupNoteRoutes(noteRouter)
}
