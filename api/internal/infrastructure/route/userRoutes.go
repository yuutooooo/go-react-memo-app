package route

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/yourusername/go-react-memo-app/internal/infrastructure/config"
	"github.com/yourusername/go-react-memo-app/internal/infrastructure/di"
	"github.com/yourusername/go-react-memo-app/internal/infrastructure/middleware"
	"github.com/yourusername/go-react-memo-app/internal/interface/controller"
)

func NewUserController() *controller.UserController {
	db, err := config.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	container := di.NewContainer(db)
	return container.UserController()
}

func SetupUserRoutes(e *echo.Group) {
	userController := NewUserController()
	e.GET("", userController.GetAllUser)
	e.POST("/signup", userController.Signup)
	e.POST("/signin", userController.Signin)
	e.POST("/logout", userController.Logout)

	// 認証が必要なルート
	auth := e.Group("")
	auth.Use(middleware.AuthMiddleware)
	auth.GET("/index", userController.Index)
}
