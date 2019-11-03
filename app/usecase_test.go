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

type mockRepo struct{}

func (r *mockRepo) FetchIndex(context.Context) ([]domain.Entry, error) {
	return []domain.Entry{
		{ID: "1", Title: "a"}, {ID: "2", Title: "b"}, {ID: "3", Title: "c"},
	}, nil
}

func (r *mockRepo) Fetch(context.Context, string) (*domain.Entry, error) {
	return nil, errors.New("not implemented")
}
