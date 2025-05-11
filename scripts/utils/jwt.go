package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// 有効期限を1週間に設定
const tokenExp = 7 * 24 * time.Hour

func GenerateJWT(sub string) (string, error) {
	// 有効期限を1週間に設定
	exp := time.Now().Add(tokenExp).Unix()

	claims := jwt.MapClaims{
		"sub": sub,
		"exp": exp,
	}

	// トークン作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 署名付きトークンを生成
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
