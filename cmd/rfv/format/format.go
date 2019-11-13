package format

import (
	"encoding/json"
	"io"

	"github.com/tomocy/rfv/domain"
)

type InJSON struct{}

func (p *InJSON) PrintIndex(w io.Writer, es []domain.Entry) {
	json.NewEncoder(w).Encode(es)
}

func (p *InJSON) Print(w io.Writer, e *domain.Entry) {
	json.NewEncoder(w).Encode(e)
}

func convert(e *domain.Entry) *entry {
	return &entry{
		ID: e.ID, Title: e.Title,
	}
}

type entry struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
