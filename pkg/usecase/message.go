package usecase

import (
	"go-app/pkg/domain"
)

type IMessageUsecase interface {
	SendMessageAndSaveResponse(message *domain.Message) (*domain.MessageWithResponse, error)
	GetMessagesAndResponseByUserID(userID int) ([]*domain.MessageWithResponse, error)
	UpdateMessageContent(messageID int, content string) error
	DeleteMessage(messageID int) error
}

type messageUsecase struct {
	mr domain.IMessageRepository
	cg IChatGPTAPIClient
}

func NewMessageUsecase(mr domain.IMessageRepository, cg IChatGPTAPIClient) IMessageUsecase {
	return &messageUsecase{
		mr,
		cg,
	}
}

func (u *messageUsecase) SendMessageAndSaveResponse(message *domain.Message) (*domain.MessageWithResponse, error) {
	messageID, err := u.mr.StoreMessage(message)
	if err != nil {
		return nil, err
	}
	message.ID = messageID

	res, err := u.cg.SendMessage(message)
	if err != nil {
		return nil, err
	}

	res.MessageID = messageID
	if err = u.mr.StoreResponse(res); err != nil {
		return nil, err
	}

	messageWithResponse := &domain.MessageWithResponse{
		ID:  message.ID,
		Content: message.Content,
		UserID: message.UserID,
		Response: res,
	}

	return messageWithResponse, nil
}

func (u *messageUsecase) GetMessagesAndResponseByUserID(userID int) ([]*domain.MessageWithResponse, error) {
	messages, err := u.mr.GetMessagesByUserID(userID)

	if err != nil {
		return nil, err
	}

	var messagesWithResponse []*domain.MessageWithResponse

	for _, m := range messages {
		response, err := u.mr.GetResponseByMessageID(m.ID)
		if err != nil {
			return nil, err
		}
		if response != nil {
			messageWithResponse := &domain.MessageWithResponse{
				ID: m.ID,
				Content: m.Content,
				UserID: m.UserID,
				Response: response,
			}
			messagesWithResponse = append(messagesWithResponse, messageWithResponse)
		}
	}

	return messagesWithResponse, nil
}

func (u *messageUsecase) UpdateMessageContent(messageID int, content string) error {
	if err := u.mr.UpdateMessageContent(messageID, content); err != nil {
		return err
	}

	return nil
}

func (u *messageUsecase) DeleteMessage(messageID int) error {
	if err := u.mr.DeleteMessage(messageID); err != nil {
		return err
	}

	return nil
}
