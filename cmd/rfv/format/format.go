package format

import (
	"encoding/json"
	"io"

	"github.com/tomocy/rfv/domain"
)

type InJSON struct{}

func (p *InJSON) PrintAll(w io.Writer, rfcs []*domain.RFC) error {
	converted := make([]*rfc, len(rfcs))
	for i, rfc := range rfcs {
		converted[i] = convert(rfc)
	}

	return json.NewEncoder(w).Encode(converted)
}

func (p *InJSON) Print(w io.Writer, rfc *domain.RFC) error {
	return json.NewEncoder(w).Encode(convert(rfc))
}

func convert(raw *domain.RFC) *rfc {
	return &rfc{
		ID:    raw.ID,
		Title: raw.Title,
	}
}

type rfc struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
