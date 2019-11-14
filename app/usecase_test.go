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
	expected := []*domain.RFC{
		{ID: 1, Title: "a"}, {ID: 2, Title: "b"}, {ID: 3, Title: "c"},
	}

	u := app.NewRFCUsecase(new(mockRepo))
	actual, err := u.Get(context.Background())
	if err != nil {
		t.Fatalf("unexpected error by (*RFCUsecase).Get: got %s, expected nil", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected entries by (*RFCUsecase).Get: got %v, expected %v", actual, expected)
	}
}

func TestFetch(t *testing.T) {
	expected := &domain.RFC{ID: 1, Title: "a"}

	u := app.NewRFCUsecase(new(mockRepo))
	actual, err := u.Find(context.Background(), 1)
	if err != nil {
		t.Fatalf("unexpected error by (*RFCUsecase).Find: got %s, expected nil", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected entry by (*RFCUsecase).Find: got %v, expected %v", actual, expected)
	}
}

type mockRepo struct{}

func (r *mockRepo) Get(context.Context) ([]*domain.RFC, error) {
	return []*domain.RFC{
		{ID: 1, Title: "a"}, {ID: 2, Title: "b"}, {ID: 3, Title: "c"},
	}, nil
}

func (r *mockRepo) Find(_ context.Context, id int) (*domain.RFC, error) {
	if id == 1 {
		return &domain.RFC{ID: 1, Title: "a"}, nil
	}

	return nil, errors.New("not found")
}
