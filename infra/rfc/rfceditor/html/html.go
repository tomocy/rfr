package html

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tomocy/rfv/infra/rfc"
)

type Old struct{}

func (s *Old) Scrape(doc *goquery.Document) (*rfc.RFC, error) {
	id, err := strconv.Atoi(strings.TrimPrefix(doc.Find(".newpage .grey a").First().Text(), "RFC "))
	if err != nil {
		return nil, fmt.Errorf("failed to convert id of rfc from id of string to id of int")
	}

	return &rfc.RFC{
		ID:    id,
		Title: doc.Find("span.h1").First().Text(),
	}, nil
}

type New struct{}

func (s *New) Scrape(doc *goquery.Document) (*rfc.RFC, error) {
	id, err := strconv.Atoi(strings.TrimPrefix(doc.Find("#rfcnum").Text(), "RFC "))
	if err != nil {
		return nil, fmt.Errorf("failed to convert id of rfc from id of string to id of int")
	}

	return &rfc.RFC{
		ID:    id,
		Title: doc.Find("#title").Text(),
	}, nil
}
