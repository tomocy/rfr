package rfc2

import "time"

type RFC struct {
	ID       int
	Title    string
	Authors  []*Author
	IssuedAt time.Time
}

type Author struct {
	Name         string
	Fullname     string
	Organization string
}
