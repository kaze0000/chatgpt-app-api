package infra

import (
	"database/sql"
	"go-app/pkg/domain"
)

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) domain.IMessageRepository {
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

func (r *messageRepository) GetMessagesByUserID(userID int) ([]*domain.Message, error) {
	rows, err := r.db.Query("SELECT id, content, user_id FROM messages WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := []*domain.Message{}
	for rows.Next() {
		message := &domain.Message{}
		err := rows.Scan(&message.ID, &message.Content, &message.UserID)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (r *messageRepository) GetResponseByMessageID(messageID int) (*domain.Response, error) {
	row := r.db.QueryRow("SELECT id, message_id, content FROM responses WHERE message_id = ?", messageID)

	response := &domain.Response{}
	err := row.Scan(&response.ID, &response.MessageID, &response.Content)
	if  err != nil {
		return nil, err
	}
	return response, nil
}

func (r *messageRepository) UpdateMessageContent(messageID int, content string) error {
	stmt, err := r.db.Prepare("UPDATE messages SET content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(content, messageID)
	return err
}

func (r *messageRepository) DeleteMessage(messageID int) error {
	stmt, err := r.db.Prepare("DELETE FROM messages WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(messageID)
	return err
}
