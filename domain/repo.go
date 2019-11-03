package domain

import "context"

type EntryRepo interface {
	FetchIndex(context.Context) ([]Entry, error)
	Fetch(context.Context, string) (*Entry, error)
}
