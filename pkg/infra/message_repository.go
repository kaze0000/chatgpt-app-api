package infra

import (
	"database/sql"
	"go-app/pkg/domain"
)

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) domain.MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) StoreMessage(m *domain.Message) error {
	stmt, err := r.db.Prepare("INSERT INTO messages (content, user_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(m.Content, m.UserID)

	return err
}

// chat gptからのレスポンスを保存するところ
func (r *messageRepository) StoreResponse(res *domain.Response) error {
	stmt, err := r.db.Prepare("INSERT INTO responses (message_id, content) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(res.MessageID, res.Content)

	return err
}
