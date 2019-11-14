package infra

import (
	"context"

	"github.com/tomocy/rfv/domain"
	rfcPkg "github.com/tomocy/rfv/infra/rfc"
	"github.com/tomocy/rfv/infra/rfc/rfceditor"
)

type ViaHTTP struct{}

func (r *ViaHTTP) Get(ctx context.Context) ([]*domain.RFC, error) {
	client := &rfcPkg.Client{
		Repo: &rfceditor.InXML{
			Fetcher: new(rfcPkg.ViaHTTP),
		},
	}
	got, err := client.Get(ctx)
	if err != nil {
		return nil, err
	}

	return adaptRFCs(got), nil
}

func (r *ViaHTTP) Find(ctx context.Context, id int) (*domain.RFC, error) {
	client := &rfcPkg.Client{
		Repo: &rfceditor.InXML{
			Fetcher: new(rfcPkg.ViaHTTP),
		},
	}
	got, err := client.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return adaptRFC(got), nil
}

func adaptRFCs(raw []*rfcPkg.RFC) []*domain.RFC {
	adapted := make([]*domain.RFC, len(raw))
	for i, rfc := range raw {
		adapted[i] = adaptRFC(rfc)
	}

	return adapted
}

func adaptRFC(raw *rfcPkg.RFC) *domain.RFC {
	return &domain.RFC{
		ID:    raw.ID,
		Title: raw.Title,
	}
}
