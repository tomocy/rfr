package rfc

import (
	"context"
	"io"
)

type Client struct {
	Fetcher Fetcher
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
