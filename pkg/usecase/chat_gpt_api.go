package usecase

import "go-app/pkg/domain"

type IChatGPTAPIClient interface {
	SendMessage(message *domain.Message) (*domain.Response, error)
}
