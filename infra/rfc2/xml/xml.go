package xml

import (
	"encoding/xml"
	"fmt"
	"io"
	"strconv"

	"github.com/tomocy/rfv/infra/rfc2"
	"github.com/tomocy/rfv/infra/rfc2/xml/rfc7991"
)

func NewRFC7991(src io.Reader) *RFC7991 {
	return &RFC7991{
		src: src,
	}
}

type RFC7991 struct {
	src io.Reader
}

func (d *RFC7991) Decode(dst *rfc2.RFC) error {
	rfc := new(rfc7991.RFC)
	if err := d.newXMLDecoder().Decode(rfc); err != nil {
		return err
	}
	if err := d.adapt(dst, rfc); err != nil {
		return err
	}

	return nil
}

func (d *RFC7991) newXMLDecoder() *xml.Decoder {
	ofXML := xml.NewDecoder(d.src)
	ofXML.Strict = false

	return ofXML
}

func (d *RFC7991) adapt(dst *rfc2.RFC, src *rfc7991.RFC) error {
	if idInfo, ok := src.Front.SeriesInfos.Find("RFC"); ok {
		id, err := strconv.Atoi(idInfo.Value)
		if err != nil {
			return fmt.Errorf("failed to parse id of string to id of int: %s", err)
		}
		dst.ID = id
	}

	dst.Title = src.Front.Title

	authors := make([]*rfc2.Author, len(src.Front.Authors))
	for i, author := range src.Front.Authors {
		authors[i] = &rfc2.Author{
			Name:         author.Name(),
			Fullname:     author.Fullname,
			Organization: author.Organization,
		}
	}
	dst.Authors = authors

	if date, err := src.Front.Date.Time(); err == nil {
		dst.IssuedAt = date
	}

	return nil
}
