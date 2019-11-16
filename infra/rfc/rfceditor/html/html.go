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
		ID:       id,
		Title:    doc.Find("span.h1").First().Text(),
		Sections: s.scrapeSections(doc),
	}, nil
}

func (s *Old) scrapeSections(doc *goquery.Document) []*rfc.Section {
	var secs []*rfc.Section
	doc.Find("span.h2").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimLeft(s.Children().Remove().End().Text(), ". ")
		if title == "" {
			return
		}

		secs = append(secs, &rfc.Section{
			Title: title,
		})
	})

	return secs
}

type New struct{}

func (s *New) Scrape(doc *goquery.Document) (*rfc.RFC, error) {
	id, err := strconv.Atoi(strings.TrimPrefix(doc.Find("#rfcnum").Text(), "RFC "))
	if err != nil {
		return nil, fmt.Errorf("failed to convert id of rfc from id of string to id of int")
	}

	return &rfc.RFC{
		ID:       id,
		Title:    doc.Find("#title").Text(),
		Sections: s.scrapeSections(doc),
	}, nil
}

func (s *New) scrapeSections(doc *goquery.Document) []*rfc.Section {
	var secs []*rfc.Section
	doc.Find("section[id^=section-]").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h2 a").Not(".section-number").First().Text()
		if title == "" {
			return
		}

		secs = append(secs, &rfc.Section{
			Title: title,
		})
	})

	return secs
}
