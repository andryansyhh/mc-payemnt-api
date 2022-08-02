package usecase

import "mc-payment-api/repository"

type McPaymentUsecase struct {
	repo     repository.RepoInterface
}

type UsecaseInterface interface {
	BookUsecase
	CategoryUsecase
	AuthorUsecase
}

func NewUsecase(repo repository.RepoInterface) UsecaseInterface {
	return &McPaymentUsecase{
		repo:     repo,
	}
}
