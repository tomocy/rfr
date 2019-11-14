package domain

import "context"

type RFCRepo interface {
	Get(context.Context) ([]*RFC, error)
	Find(context.Context, int) (*RFC, error)
}
