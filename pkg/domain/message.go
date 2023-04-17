package domain

type Message struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	UserID		string    `json:"user_id"`
}

type Response struct {
	ID        int    `json:"id"`
	MessageID int    `json:"message_id"`
	Content   string `json:"content"`
}

type MessageRepository interface {
	StoreMessage(m *Message) error
	StoreResponse(r *Response) error
	// TODO: 必要あれば実装する
	// GetMessages() ([]*Message, error)
	// GetResponses(messageID int) ([]*Response, error)
}
