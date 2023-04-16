package usecase

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"go-app/pkg/domain"
)


func hashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4)

	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func StoreUser(user *domain.User, repo domain.UserRepository) error {
	hashPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashPassword

	return repo.Store(user)
}

func AuthenticateUser(email, password string, repo domain.UserRepository, secretKey string) (string, error) {
	user, err := repo.FindByEmail(email)

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

	token, err := GenerateJWT(user, secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
