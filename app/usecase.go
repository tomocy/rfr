package app

import (
	"context"

	"github.com/tomocy/rfv/domain"
)

func NewEntryUsecase(repo domain.EntryRepo) *EntryUsecase {
	return &EntryUsecase{
		repo: repo,
	}
}

type EntryUsecase struct {
	repo domain.EntryRepo
}

func (u *EntryUsecase) FetchIndex(ctx context.Context) ([]domain.Entry, error) {
	return u.repo.FetchIndex(ctx)
}

func (u *EntryUsecase) Fetch(ctx context.Context, id string) (*domain.Entry, error) {
	return u.repo.Fetch(ctx, id)
}
