package domain

type Profile struct {
	ID        int    `json:"id"`
	UserID int `json:"user_id"`
	Hobby   string `json:"hobby"`
}

type IProfileRepository interface {
	CreateProfile(profile *Profile) (*Profile, error)
}
