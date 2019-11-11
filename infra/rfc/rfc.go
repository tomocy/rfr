package rfc

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
