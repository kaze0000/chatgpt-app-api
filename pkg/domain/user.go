package domain

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	Store(u *User) error
	FindByEmail(email string) (*User, error)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
