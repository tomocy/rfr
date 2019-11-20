package rfc2

type RFC struct {
	ID      int
	Title   string
	Authors []*Author
}

type Author struct {
	Name         string
	Fullname     string
	Organization string
}
