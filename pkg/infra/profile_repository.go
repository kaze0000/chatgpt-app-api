package infra

import (
	"database/sql"
	"go-app/pkg/domain"
)

type profileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) domain.IProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) CreateProfile(p *domain.Profile) (*domain.Profile, error) {
	stmt, err := r.db.Prepare("INSERT INTO profiles (user_id, hobby) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.UserID, p.Hobby)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	p.ID = int(id)
	if err != nil {
		return nil, err
	}

	return p, err
}

func (r *profileRepository) GetProfileByUserID(userID int) (*domain.Profile, error) {
	row := r.db.QueryRow("SELECT id, user_id, hobby FROM profiles WHERE user_id = ?", userID)

	profile := &domain.Profile{}
	err := row.Scan(&profile.ID, &profile.UserID, &profile.Hobby)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
