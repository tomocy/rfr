package mode

import (
	"io"

	"github.com/tomocy/chi"
	"github.com/tomocy/rfv/app"
	"github.com/tomocy/rfv/domain"
)

type OnHTTP struct {
	addr    string
	router  chi.Router
	usecase app.EntryUsecase
	printer Printer
}

type Printer interface {
	PrintIndex(io.Writer, []domain.Entry)
	Print(io.Writer, *domain.Entry)
}
