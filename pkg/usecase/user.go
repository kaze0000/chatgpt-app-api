package usecase

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"go-app/pkg/domain"
)

type IUserUsecase interface {
	StoreUser(user *domain.User) error
	AuthenticateUser(email, password string) (string, *http.Cookie, error)
}

type userUsecase struct {
	repo			domain.IUserRepository
	secretKey	string
}

func NewUserUsecase(repo domain.IUserRepository, secretKey string) IUserUsecase {
	return &userUsecase{repo: repo, secretKey: secretKey}
}

func hashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4)

	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func (u *userUsecase) StoreUser(user *domain.User) error {
	hashPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashPassword

	return u.repo.Store(user)
}

func (u *userUsecase) AuthenticateUser(email, password string) (string, *http.Cookie, error) {
	user, err := u.repo.FindByEmail(email)

	if err != nil {
		return "", nil, err
	}
	if user == nil {
		return "", nil, fmt.Errorf("user not found")
	}

	// パスワードの比較
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", nil, err
	}

	// トークンの生成
	token, err := GenerateJWT(user, u.secretKey)
	if err != nil {
		return "", nil, err
	}

	// tokenをset-cookieに入れる
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode

	return token, cookie, nil
}
