package controller

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yourusername/go-react-memo-app/internal/interface/dto"
	"github.com/yourusername/go-react-memo-app/internal/usecase"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (c *UserController) GetAllUser(ctx echo.Context) error {
	users, err := c.userUsecase.GetAllUser()
	if err != nil {
		response := dto.FailResponse("ユーザー一覧の取得に失敗しました", err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	usersResponse := dto.CreateUserResponseMany(users)
	response := dto.SuccessResponse(usersResponse, "ユーザー一覧を取得しました")
	return ctx.JSON(http.StatusOK, response)
}

func (c *UserController) Signup(ctx echo.Context) error {
	req, err := dto.CreateUserRequest(ctx)
	if err != nil {
		response := dto.ErrorResponse("リクエストの形式が正しくありません", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	if err := c.userUsecase.CheckEmail(req.Email); err != nil {
		response := dto.ErrorResponse("このメールアドレスは既に使用されています", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	user, token, err := c.userUsecase.CreateUser(req)
	if err != nil {
		response := dto.FailResponse("ユーザー登録に失敗しました", err)
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	// JWTトークンをHTTP-Only Secure Cookieとして設定
	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = token
	cookie.HttpOnly = true                    // JavaScriptからアクセス不可
	cookie.Secure = true                      // HTTPSのみで送信
	cookie.SameSite = http.SameSiteStrictMode // CSRF対策
	cookie.Path = "/"
	// 有効期限はJWTトークンと同じに設定（例：24時間）
	cookie.MaxAge = 24 * 60 * 60 // 24時間（秒単位）
	ctx.SetCookie(cookie)

	userResponse := dto.CreateUserResponseSingle(*user, token)
	response := dto.SuccessResponse(userResponse, "ユーザー登録が完了しました")
	return ctx.JSON(http.StatusCreated, response)
}

func (c *UserController) Signin(ctx echo.Context) error {
	req, err := dto.SigninUserRequest(ctx)
	if err != nil {
		response := dto.ErrorResponse("リクエストの形式が正しくありません", err)
		return ctx.JSON(http.StatusBadRequest, response)
	}

	user, token, err := c.userUsecase.Login(req)
	if err != nil {
		response := dto.ErrorResponse("ログインに失敗しました。メールアドレスまたはパスワードが正しくありません", err)
		return ctx.JSON(http.StatusUnauthorized, response)
	}

	// JWTトークンをHTTP-Only Secure Cookieとして設定
	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = token
	cookie.HttpOnly = true                    // JavaScriptからアクセス不可
	cookie.Secure = true                      // HTTPSのみで送信
	cookie.SameSite = http.SameSiteStrictMode // CSRF対策
	cookie.Path = "/"
	// 有効期限はJWTトークンと同じに設定（例：24時間）
	cookie.MaxAge = 24 * 60 * 60 // 24時間（秒単位）
	ctx.SetCookie(cookie)

	userResponse := dto.CreateUserResponseSingle(*user, token)
	response := dto.SuccessResponse(userResponse, "ログインに成功しました")
	return ctx.JSON(http.StatusOK, response)
}

func (c *UserController) Index(ctx echo.Context) error {
	userID := ctx.Get("user_id").(string)
	if userID == "" {
		return ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse("ユーザーが見つかりません", errors.New("ユーザーが見つかりません")))
	}
	user, err := c.userUsecase.GetUserById(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse("ユーザーの取得に失敗しました", err))
	}
	folders, notes, err := c.userUsecase.GetAllFolderAndNoteByUserID(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse("フォルダーとノートの取得に失敗しました", err))
	}
	folderAndNoteTree, err := dto.CreateFolderNoteTree(folders, notes)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse("フォルダーとノートのツリーの取得に失敗しました", err))
	}
	response := dto.CreateUserIndexResponse(*user, folderAndNoteTree)
	return ctx.JSON(http.StatusOK, response)
}

func (c *UserController) Logout(ctx echo.Context) error {
	// クッキーを削除するために、同じ名前で有効期限切れのクッキーを設定
	cookie := new(http.Cookie)
	cookie.Name = "auth_token"
	cookie.Value = ""
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Path = "/"
	cookie.MaxAge = -1 // 有効期限切れに設定
	ctx.SetCookie(cookie)

	response := dto.SuccessResponse(nil, "ログアウトに成功しました")
	return ctx.JSON(http.StatusOK, response)
}
