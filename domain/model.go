package domain

type RFC struct {
	ID       int
	Title    string
	Sections []*Section
}

type Section struct {
	Title    string
	Body     string
	Sectinos []*Section
}
