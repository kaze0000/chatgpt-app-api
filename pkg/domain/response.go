package domain

type Response struct {
	ID        int    `json:"id"`
	MessageID int `json:"message_id"`
	Content   string `json:"content"`
}

