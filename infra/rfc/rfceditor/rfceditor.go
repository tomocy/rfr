package rfceditor

import (
	"context"
	"encoding/xml"
	"io"

	"github.com/tomocy/rfv/infra/rfc"
	rfcxml "github.com/tomocy/rfv/infra/rfc/rfceditor/xml"
)

type InXML struct {
	Fetcher Fetcher
}

func (r *InXML) Get(ctx context.Context) ([]*rfc.RFC, error) {
	fetched, err := r.Fetcher.Fetch(ctx, "//www.rfc-editor.org/rfc-index.xml")
	if err != nil {
		return nil, err
	}
	defer fetched.Close()

	index := new(rfcxml.Index)
	if err := xml.NewDecoder(fetched).Decode(index); err != nil {
		return nil, err
	}

	return r.convertIndex(index), nil
}

func (r *InXML) convertIndex(raw *rfcxml.Index) []*rfc.RFC {
	converted := make([]*rfc.RFC, len(raw.Entries))
	for i, entry := range raw.Entries {
		converted[i] = r.convertEntry(entry)
	}

	return converted
}

func (r *InXML) convertEntry(raw *rfcxml.Entry) *rfc.RFC {
	return &rfc.RFC{
		ID:    raw.ID,
		Title: raw.Title,
	}
}

func (r *InXML) Find(ctx context.Context, id int) (*rfc.RFC, error) {
	fetched, err := r.Fetcher.Fetch(ctx, "single.xml")
	if err != nil {
		return nil, err
	}
	defer fetched.Close()

	rfc := new(rfcxml.RFC)
	if err := xml.NewDecoder(fetched).Decode(rfc); err != nil {
		return nil, err
	}

	return r.convertRFC(rfc), nil
}

func (r *InXML) convertRFC(raw *rfcxml.RFC) *rfc.RFC {
	return &rfc.RFC{
		Title: raw.Front.Title,
	}
}

type Fetcher interface {
	Fetch(context.Context, string) (io.ReadCloser, error)
}
