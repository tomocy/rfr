package rfc

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/tomocy/rfv/infra/rfc/text"
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
	Authors  []string
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

type InText struct {
	URI     URI
	Fetcher Fetcher
}

func (r *InText) Get(ctx context.Context) ([]*RFC, error) {
	fetched, err := r.Fetcher.Fetch(ctx, r.URI.OfIndex("txt"))
	if err != nil {
		return nil, err
	}
	defer fetched.Close()

	formatted, err := r.reformatForIndex(fetched)
	if err != nil {
		return nil, err
	}

	index := new(text.Index)
	if _, err := fmt.Fscan(formatted, index); err != nil {
		return nil, err
	}

	rfcs := make([]*RFC, len(index.Lines))
	for i, line := range index.Lines {
		rfcs[i] = &RFC{
			ID:    line.ID,
			Title: string(line.Title),
		}
	}

	return rfcs, nil
}

func (r *InText) reformatForIndex(target io.Reader) (io.Reader, error) {
	target, err := r.readerFromIndex(target)
	if err != nil {
		return nil, err
	}

	var elems, lines []string
	scanner := bufio.NewScanner(target)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			lines = append(lines, strings.Join(elems, " "))
			elems = nil
			continue
		}

		elems = append(elems, strings.Trim(text, " "))
	}
	lines = append(lines, strings.Join(elems, " "))
	elems = nil

	return strings.NewReader(strings.Join(lines, "\n")), nil
}

func (r *InText) readerFromIndex(src io.Reader) (io.Reader, error) {
	read, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, err
	}
	target := string(read)

	begin := strings.Index(target, "0001")
	if begin < 0 {
		return nil, errors.New("the start position to parse index is not found")
	}

	return strings.NewReader(target[begin:]), nil
}

type Fetcher interface {
	Fetch(context.Context, string) (io.ReadCloser, error)
}

type URI interface {
	OfIndex(string) string
	Of(string, int) string
}

type RFCEditor struct{}

func (r *RFCEditor) OfIndex(ext string) string {
	switch ext {
	case "txt":
		return r.endpoint("rfc-index.txt")
	default:
		return ""
	}
}

func (r *RFCEditor) Of(ext string, id int) string {
	switch ext {
	case "txt":
		return r.endpoint("rfc", fmt.Sprintf("rfc%d.txt", id))
	default:
		return ""
	}
}

func (r *RFCEditor) endpoint(ps ...string) string {
	return "//" + filepath.Join(append([]string{"www.rfc-editor.org"}, ps...)...)
}
