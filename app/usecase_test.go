package app_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/tomocy/rfv/app"
	"github.com/tomocy/rfv/domain"
)

func TestFetchIndex(t *testing.T) {
	expected := []domain.Entry{
		{ID: "1", Title: "a"}, {ID: "2", Title: "b"}, {ID: "3", Title: "c"},
	}

	u := app.NewEntryUsecase(new(mockRepo))
	actual, err := u.FetchIndex(context.Background())
	if err != nil {
		t.Fatalf("unexpected error by (*EntryUsecase).FetchIndex: got %s, expected nil", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected entries by (*EntryUsecase).FetchIndex: got %v, expected %v", actual, expected)
	}
}

func TestFetch(t *testing.T) {
	expected := &domain.Entry{ID: "1", Title: "a"}

	u := app.NewEntryUsecase(new(mockRepo))
	actual, err := u.Fetch(context.Background(), "1")
	if err != nil {
		t.Fatalf("unexpected error by (*EntryUsecase).Fetch: got %s, expected nil", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected entry by (*EntryUsecase).Fetch: got %v, expected %v", actual, expected)
	}
}

type mockRepo struct{}

func (r *mockRepo) FetchIndex(context.Context) ([]domain.Entry, error) {
	return []domain.Entry{
		{ID: "1", Title: "a"}, {ID: "2", Title: "b"}, {ID: "3", Title: "c"},
	}, nil
}

func (r *mockRepo) Fetch(_ context.Context, id string) (*domain.Entry, error) {
	if id == "1" {
		return &domain.Entry{ID: "1", Title: "a"}, nil
	}

	return nil, errors.New("not found")
}
