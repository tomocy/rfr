package rfc7991

import (
	"fmt"
	"strings"
)

type RFC struct {
	Front Front `xml:"front"`
}

type Front struct {
	Title       string      `xml:"title"`
	SeriesInfos SeriesInfos `xml:"seriesInfo"`
	Authors     []*Author   `xml:"author"`
}

type SeriesInfos []*SeriesInfo

func (is SeriesInfos) Find(name string) (*SeriesInfo, bool) {
	for _, info := range is {
		if info.Name == name {
			return info, true
		}
	}

	return nil, false
}

type SeriesInfo struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Author struct {
	Initials     string `xml:"initials,attr"`
	Surname      string `xml:"surname,attr"`
	Fullname     string `xml:"fullname,attr"`
	Organization string `xml:"organization"`
}

func (a *Author) Name() string {
	initials := strings.Trim(a.Initials, " ")
	if !strings.HasSuffix(a.Initials, ".") {
		initials += "."
	}

	return fmt.Sprintf("%s %s", initials, a.Surname)
}
