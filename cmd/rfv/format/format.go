package format

import (
	"encoding/json"
	"io"
	"log"

	"github.com/tomocy/rfv/domain"
)

type InJSON struct{}

func (p *InJSON) PrintIndex(w io.Writer, es []domain.Entry) {
	converted := make([]entry, len(es))
	for i, e := range es {
		converted[i] = *convert(&e)
	}

	if err := json.NewEncoder(w).Encode(converted); err != nil {
		log.Println(err)
	}
}

func (p *InJSON) Print(w io.Writer, e *domain.Entry) {
	json.NewEncoder(w).Encode(convert(e))
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
