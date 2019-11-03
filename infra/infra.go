package infra

import (
	"github.com/tomocy/rff"
	"github.com/tomocy/rfv/domain"
)

type ViaHTTP struct{}

type index rff.Index

type rfc rff.RFC

func (r *rfc) adapt() *domain.Entry {
	return &domain.Entry{
		Title: r.Front.Title,
	}
}
