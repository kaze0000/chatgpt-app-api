package usecase

import "go-app/pkg/domain"

type ChatGPTAPI interface {
	SendMessage(message *domain.Message) (*domain.Response, error)
}
