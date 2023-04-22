package infra

import (
	"database/sql"
	"go-app/pkg/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Store(u *domain.User) error {
	stmt, err := r.db.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Name, u.Email, u.Password)
	return err
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	row := r.db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email)

	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
