package infra

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/tomocy/rfv/domain"
	rfcPkg "github.com/tomocy/rfv/infra/rfc"
	"github.com/tomocy/rfv/infra/rfc/rfceditor"
)

func NewViaHTTP() *ViaHTTP {
	return &ViaHTTP{
		client: &rfcPkg.Client{
			Repo: &rfceditor.InXML{
				Fetcher: new(rfcPkg.ViaHTTP),
			},
		},
	}
}

type ViaHTTP struct {
	client *rfcPkg.Client
}

func (r *ViaHTTP) FetchIndex(ctx context.Context) ([]domain.Entry, error) {
	got, err := r.client.Get(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rfcs(got).adapt(), nil
}

func (r *ViaHTTP) Fetch(ctx context.Context, id string) (*domain.Entry, error) {
	idNum, err := strconv.Atoi(strings.TrimLeft(id, "rfc"))
	if err != nil {
		return nil, fmt.Errorf("failed to convert id of string to id of int")
	}

	got, err := r.client.Find(ctx, idNum)
	if err != nil {
		return nil, err
	}

	converted := rfc(*got)
	return converted.adapt(), nil
}

type rfcs []*rfcPkg.RFC

func (rs rfcs) adapt() []domain.Entry {
	adapted := make([]domain.Entry, len(rs))
	for i, r := range rs {
		converted := rfc(*r)
		adapted[i] = *converted.adapt()
	}

	return adapted
}

type rfc rfcPkg.RFC

func (r *rfc) adapt() *domain.Entry {
	return &domain.Entry{
		ID:    r.ID,
		Title: r.Title,
	}
}
