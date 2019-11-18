package domain

type RFC struct {
	ID    int
	Title string
}

type Section struct {
	Title    string
	Body     string
	Sectinos []*Section
}
