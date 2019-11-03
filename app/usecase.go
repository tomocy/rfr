package app

import (
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
