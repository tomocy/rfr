package format

import (
	"encoding/json"
	"io"

	"github.com/tomocy/rfv/domain"
)

type InJSON struct{}

func (p *InJSON) PrintIndex(w io.Writer, es []domain.Entry) {
	json.NewEncoder(w).Encode(es)
}
