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
		ID:       raw.ID,
		Title:    raw.Title,
		Sections: convertSections(raw.Sections),
	}
}

func convertSections(raw []*domain.Section) []*section {
	converted := make([]*section, len(raw))
	for i, sec := range raw {
		converted[i] = convertSection(sec)
	}

	return converted
}

func convertSection(raw *domain.Section) *section {
	return &section{
		Title:    raw.Title,
		Body:     raw.Body,
		Sections: convertSections(raw.Sectinos),
	}
}

type rfc struct {
	ID       int        `json:"id"`
	Title    string     `json:"title"`
	Sections []*section `json:"sections"`
}

type section struct {
	Title    string     `json:"title"`
	Body     string     `json:"body"`
	Sections []*section `json:"sections"`
}
