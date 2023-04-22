package usecase

import "go-app/pkg/domain"

type IProfileUsecase interface {
	CreateProfile(profile *domain.Profile) (*domain.Profile, error)
}

type profileUsecase struct {
	pr domain.IProfileRepository
}

func NewProfileUsecase(pr domain.IProfileRepository) IProfileUsecase {
	return &profileUsecase{pr}
}

func (u *profileUsecase) CreateProfile(profile *domain.Profile) (*domain.Profile, error) {
	return u.pr.CreateProfile(profile)
}
