package rfc7991

import (
	"fmt"
	"strings"
	"time"
)

type RFC struct {
	Front Front `xml:"front"`
}

type Front struct {
	Title       string      `xml:"title"`
	SeriesInfos SeriesInfos `xml:"seriesInfo"`
	Authors     []*Author   `xml:"author"`
	Date        Date        `xml:"date"`
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

type Date struct {
	Day   string `xml:"day,attr"`
	Month string `xml:"month,attr"`
	Year  string `xml:"year,attr"`
}

func (d *Date) Time() (time.Time, error) {
	var formatElems, targetElems []string
	if d.Day != "" {
		formatElems, targetElems = append(formatElems, "02"), append(targetElems, d.Day)
	}
	if d.Month != "" {
		if strings.Contains(
			"1 2 3 4 5 6 7 8 9 10 11 12",
			strings.TrimLeft(d.Month, "0"),
		) {
			formatElems, targetElems = append(formatElems, "01"), append(targetElems, d.Month)
		} else {
			formatElems, targetElems = append(formatElems, "January"), append(targetElems, d.Month)
		}
	}
	if d.Year != "" {
		formatElems, targetElems = append(formatElems, "2006"), append(targetElems, d.Year)
	}

	return time.Parse(
		strings.Join(formatElems, " "),
		strings.Join(targetElems, " "),
	)
}
