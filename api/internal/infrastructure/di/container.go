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
	folderRepository repository.FolderRepository

	// サービス
	userService service.UserService
	folderService service.FolderService

	// ユースケース
	userUsecase usecase.UserUsecase
	folderUsecase usecase.FolderUsecase

	// コントローラー
	userController *controller.UserController
	folderController *controller.FolderController
}

func 	NewContainer(db *gorm.DB) *Container {
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
	c.folderRepository = persistence.NewFolderRepository(c.db)

	// サービスの初期化
	c.userService = service.NewUserService(c.userRepository)
	c.folderService = service.NewFolderService(c.folderRepository)

	// ユースケースの初期化
	c.userUsecase = usecase.NewUserUsecase(c.userService, c.userRepository)
	c.folderUsecase = usecase.NewFolderUsecase(c.folderService)

	// コントローラーの初期化
	c.userController = controller.NewUserController(c.userUsecase)
	c.folderController = controller.NewFolderController(c.folderUsecase)
}

// UserControllerインスタンスを返す
func (c *Container) UserController() *controller.UserController {
	return c.userController
}

// FolderControllerインスタンスを返す
func (c *Container) FolderController() *controller.FolderController {
	return c.folderController
}
