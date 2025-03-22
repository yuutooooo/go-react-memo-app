package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tokenString string

		// まずクッキーからトークンを取得を試みる
		cookie, err := c.Cookie("auth_token")
		if err == nil && cookie.Value != "" {
			tokenString = cookie.Value
		} else {
			// クッキーがない場合はAuthorizationヘッダーを確認
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, "認証情報がありません")
			}

			// Bearer トークンの抽出
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, "認証形式が不正です")
			}
			tokenString = parts[1]
		}

		// JWT シークレットキーの取得
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			log.Println("警告: JWT_SECRET環境変数が設定されていません。デフォルト値を使用します。")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			log.Printf("トークン検証エラー: %v", err)
			return c.JSON(http.StatusUnauthorized, "無効なトークンです")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// ユーザーIDの検証
			if claims["sub"] == nil {
				return c.JSON(http.StatusUnauthorized, "トークンにユーザーIDがありません")
			}
			c.Set("user_id", claims["sub"])
			userid := c.Get("user_id").(string)
			println("user id", userid)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, "無効なトークンです")
	}
}
