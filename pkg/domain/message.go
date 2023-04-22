package domain

type Message struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	UserID		int    `json:"user_id"`
}

type MessageWithResponse struct {
		ID int `json:"id"`
		Content string `json:"content"`
		UserID int `json:"user_id"`
		Response *Response `json:"response"`
}

type IMessageRepository interface {
	StoreMessage(m *Message) error
	StoreResponse(r *Response) error
	// TODO: 必要あれば実装する
	// GetMessages() ([]*Message, error)
	// GetResponses(messageID int) ([]*Response, error)
	GetMessagesByUserID(userID int) ([]*Message, error)
	GetResponseByMessageID(messageID int) (*Response, error)
	UpdateMessageContent(messageID int, content string) error
	DeleteMessage(messageID int) error
}
