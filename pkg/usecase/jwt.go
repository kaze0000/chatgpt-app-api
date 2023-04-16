package usecase

import (
	"fmt"
	"go-app/pkg/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

const jwtExpireDuration = 24 * 60 * 60 * time.Second

type JWTClaims struct {
	UserID int `json:"user_id"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

func GenerateJWT(user *domain.User, secretKey string) (string, error) {
	claims := JWTClaims{
		UserID: user.ID,
		Name:  user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtExpireDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return signedToken, nil
}
