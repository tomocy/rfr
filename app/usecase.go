package app

import (
	"context"

	"github.com/tomocy/rfv/domain"
)

func NewRFCUsecase(repo domain.RFCRepo) *RFCUsecase {
	return &RFCUsecase{
		repo: repo,
	}
}

type RFCUsecase struct {
	repo domain.RFCRepo
}

func (u *RFCUsecase) Get(ctx context.Context) ([]*domain.RFC, error) {
	return u.repo.Get(ctx)
}

func (u *RFCUsecase) Find(ctx context.Context, id int) (*domain.RFC, error) {
	return u.repo.Find(ctx, id)
}
