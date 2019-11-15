package rfceditor

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tomocy/rfv/infra/rfc"
	"github.com/tomocy/rfv/infra/rfc/rfceditor/html"
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
	id, _ := strconv.Atoi(strings.TrimLeft(strings.ToLower(raw.ID), "rfc"))
	return &rfc.RFC{
		ID:    id,
		Title: raw.Title,
	}
}

func (r *InXML) Find(ctx context.Context, id int) (*rfc.RFC, error) {
	fetched, err := r.Fetcher.Fetch(ctx, fmt.Sprintf("//www.rfc-editor.org/rfc/rfc%d.xml", id))
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

type InHTML struct {
	Fetcher Fetcher
}

func (r *InHTML) Find(ctx context.Context, id int) (*rfc.RFC, error) {
	fetched, err := r.Fetcher.Fetch(ctx, fmt.Sprintf("//www.rfc-editor.org/rfc/rfc%d.html", id))
	if err != nil {
		return nil, err
	}
	defer fetched.Close()

	return r.scrapeRFC(fetched)
}

func (r *InHTML) scrapeRFC(src io.Reader) (*rfc.RFC, error) {
	doc, err := goquery.NewDocumentFromReader(src)
	if err != nil {
		return nil, err
	}

	if isOld := doc.Find("#name-table-of-contents").Length() < 1; isOld {
		return new(html.Old).Scrape(doc)
	}

	return new(html.New).Scrape(doc)
}

type Fetcher interface {
	Fetch(context.Context, string) (io.ReadCloser, error)
}
