package usecase

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"go-app/pkg/domain"
)

type IUserUsecase interface {
	StoreUser(user *domain.User) error
	AuthenticateUser(email, password string) (string, error)
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

func (u *userUsecase) AuthenticateUser(email, password string) (string, error) {
	user, err := u.repo.FindByEmail(email)

	if err != nil {
		return "", err
	}
	if user == nil {
		return "", fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := GenerateJWT(user, u.secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
