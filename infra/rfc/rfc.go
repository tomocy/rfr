package rfc

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
)

type Client struct {
	Fetcher Fetcher
}

func (c *Client) FetchIndex(ctx context.Context) (*Index, error) {
	fetched, err := c.Fetcher.Fetch(ctx, "//www.rfc-editor.org/rfc-index.xml")
	if err != nil {
		return nil, err
	}
	defer fetched.Close()

	index := new(Index)
	if err := xml.NewDecoder(fetched).Decode(index); err != nil {
		return nil, fmt.Errorf("failed to decode: %s", err)
	}

	return index, nil
}

type Fetcher interface {
	Fetch(context.Context, string) (io.ReadCloser, error)
}

type Index struct {
	RFCs []Entry `xml:"rfc-entry"`
}

type Entry struct {
	ID    string `xml:"doc-id"`
	Title string `xml:"title"`
}

type RFC struct {
	Front Front `xml:"front"`
}

type Front struct {
	Title string `xml:"title"`
}
