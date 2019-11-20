package rfc7991

type RFC struct {
	Front Front `xml:"front"`
}

type Front struct {
	Title       string      `xml:"title"`
	SeriesInfos SeriesInfos `xml:"seriesInfo"`
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
