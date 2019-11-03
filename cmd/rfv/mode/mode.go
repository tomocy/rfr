package mode

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

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

func (r *OnHTTP) Run() error {
	r.register()

	if err := http.ListenAndServe(r.addr, r.router); err != nil {
		return fmt.Errorf("failed to listen and serve: %s", err)
	}

	return nil
}

func (r *OnHTTP) register() {
	r.router.Get("/", r.fetchIndex)
	r.router.Get("/{id}", r.fetch)
}

func (r *OnHTTP) fetchIndex(w http.ResponseWriter, _ *http.Request) {
	idx, err := r.usecase.FetchIndex(context.Background())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.printer.PrintIndex(w, idx)
}

func (r *OnHTTP) fetch(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	e, err := r.usecase.Fetch(context.Background(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.printer.Print(w, e)
}

type Printer interface {
	PrintIndex(io.Writer, []domain.Entry)
	Print(io.Writer, *domain.Entry)
}