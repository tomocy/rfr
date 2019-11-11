package infra

import (
	"context"

	"github.com/tomocy/rfv/domain"
	rfcPkg "github.com/tomocy/rfv/infra/rfc"
)

type ViaHTTP struct{}

func (r *ViaHTTP) FetchIndex(ctx context.Context) ([]domain.Entry, error) {
	idx, err := rfcPkg.HTTPClient.FetchIndex(ctx)
	if err != nil {
		return nil, err
	}

	converted := index(*idx)
	return converted.adapt(), nil
}

func (r *ViaHTTP) Fetch(ctx context.Context, id string) (*domain.Entry, error) {
	e, err := rfcPkg.HTTPClient.Fetch(ctx, id)
	if err != nil {
		return nil, err
	}

	converted := rfc(*e)
	return converted.adapt(), nil
}

type index rfcPkg.Index

func (i *index) adapt() []domain.Entry {
	es := make([]domain.Entry, len(i.RFCs))
	for i, rfc := range i.RFCs {
		entry := entry(rfc)
		es[i] = *entry.adapt()
	}

	return es
}

type entry rfcPkg.Entry

func (e *entry) adapt() *domain.Entry {
	return &domain.Entry{
		ID: e.ID, Title: e.Title,
	}
}

type rfc rfcPkg.RFC

func (r *rfc) adapt() *domain.Entry {
	return &domain.Entry{
		Title: r.Front.Title,
	}
}
