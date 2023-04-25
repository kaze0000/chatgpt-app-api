package graph

import "go-app/pkg/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	ProfileUsecase usecase.IProfileUsecase
}

func NewResolver(pu usecase.IProfileUsecase) *Resolver {
	return &Resolver{pu}
}
