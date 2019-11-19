package rfc

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Repo Repo
}

func (c *Client) Get(ctx context.Context) ([]*RFC, error) {
	return c.Repo.Get(ctx)
}

func (c *Client) Find(ctx context.Context, id int) (*RFC, error) {
	return c.Repo.Find(ctx, id)
}

type Repo interface {
	Get(context.Context) ([]*RFC, error)
	Find(context.Context, int) (*RFC, error)
}

type RFC struct {
	ID       int
	Title    string
	Sections []*Section
}

type Section struct {
	Title    string
	Body     string
	Sections []*Section
}

type ViaHTTP struct {
	IsSecure bool
}

func (f *ViaHTTP) Fetch(ctx context.Context, uri string) (io.ReadCloser, error) {
	uri = f.compensate(uri)

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (f *ViaHTTP) compensate(uri string) string {
	scheme := "https"
	if !f.IsSecure {
		scheme = "http"
	}

	return fmt.Sprintf("%s:%s", scheme, uri)
}

type URI interface {
	OfIndex(string) string
	Of(string, int) string
}
