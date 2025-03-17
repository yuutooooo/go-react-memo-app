package di

import (
	"github.com/yourusername/go-react-memo-app/internal/domain/repository"
	"github.com/yourusername/go-react-memo-app/internal/domain/service"
	"github.com/yourusername/go-react-memo-app/internal/infrastructure/persistence"
	"github.com/yourusername/go-react-memo-app/internal/interface/controller"
	"github.com/yourusername/go-react-memo-app/internal/usecase"
	"gorm.io/gorm"
)

// Container は依存性を管理するコンテナ
type Container struct {
	db *gorm.DB

	// リポジトリ
	userRepository repository.UserRepository

	// サービス
	userService service.UserService

	// ユースケース
	userUsecase usecase.UserUsecase

	// コントローラー
	userController *controller.UserController
}

func NewContainer(db *gorm.DB) *Container {
	c := &Container{
		db: db,
	}
	c.initialize()
	return c
}

// initialize はコンテナの依存関係を初期化する
func (c *Container) initialize() {
	// リポジトリの初期化
	c.userRepository = persistence.NewUserRepository(c.db)

	// サービスの初期化
	c.userService = service.NewUserService(c.userRepository)

	// ユースケースの初期化
	c.userUsecase = usecase.NewUserUsecase(c.userService, c.userRepository)

	// コントローラーの初期化
	c.userController = controller.NewUserController(c.userUsecase)
}

// UserControllerインスタンスを返す
func (c *Container) UserController() *controller.UserController {
	return c.userController
}
