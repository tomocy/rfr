package infra

import (
	"github.com/tomocy/rff"
	"github.com/tomocy/rfv/domain"
)

type ViaHTTP struct{}

type index rff.Index

func (i *index) adapt() []domain.Entry {
	es := make([]domain.Entry, len(i.RFCs))
	for i, rfc := range i.RFCs {
		entry := entry(rfc)
		es[i] = *entry.adapt()
	}

	return es
}

type entry rff.Entry

func (e *entry) adapt() *domain.Entry {
	return &domain.Entry{
		ID: e.ID, Title: e.Title,
	}
}

type rfc rff.RFC

func (r *rfc) adapt() *domain.Entry {
	return &domain.Entry{
		Title: r.Front.Title,
	}
}
