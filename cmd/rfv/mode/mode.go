package mode

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/tomocy/rfv/app"
	"github.com/tomocy/rfv/domain"
	"github.com/tomocy/rfv/infra"
	pb "github.com/tomocy/rfv/infra/rpc/rfv"
	"google.golang.org/grpc"
)

func NewOnHTTP(addr string, printer Printer) *OnHTTP {
	return &OnHTTP{
		addr:    addr,
		router:  chi.NewRouter(),
		usecase: *app.NewEntryUsecase(new(infra.ViaHTTP)),
		printer: printer,
	}
}

type OnHTTP struct {
	addr    string
	router  chi.Router
	usecase app.EntryUsecase
	printer Printer
}

func (r *OnHTTP) Run() error {
	r.register()

	r.logf("listen and serve on %s", r.addr)
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.printer.PrintIndex(w, idx)
}

func (r *OnHTTP) fetch(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	e, err := r.usecase.Fetch(context.Background(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.printer.Print(w, e)
}

func (r *OnHTTP) logf(format string, as ...interface{}) {
	logf(fmt.Sprintf("http: %s", format), as...)
}

func NewOnGRPC(addr string) *OnGRPC {
	return &OnGRPC{
		addr:    addr,
		server:  *grpc.NewServer(),
		usecase: *app.NewEntryUsecase(new(infra.ViaHTTP)),
	}
}

type OnGRPC struct {
	pb.UnimplementedEntryRepoServer
	addr    string
	server  grpc.Server
	usecase app.EntryUsecase
}

func (r *OnGRPC) Run() error {
	pb.RegisterEntryRepoServer(&r.server, r)

	r.logf("listen and serve on %s", r.addr)
	if err := r.listenAndServe(); err != nil {
		return fmt.Errorf("failed to listen and serve: %s", err)
	}

	return nil
}

func (r *OnGRPC) listenAndServe() error {
	l, err := net.Listen("tcp", r.addr)
	if err != nil {
		return err
	}

	return r.server.Serve(l)
}

func (r *OnGRPC) FetchIndex(ctx context.Context, req *empty.Empty) (*pb.Entries, error) {
	idx, err := r.usecase.FetchIndex(ctx)
	if err != nil {
		return nil, err
	}

	converteds := make([]*pb.Entry, len(idx))
	for i, e := range idx {
		converteds[i] = r.convert(&e)
	}

	return &pb.Entries{
		Entries: converteds,
	}, nil
}

func (r *OnGRPC) Fetch(ctx context.Context, req *pb.FetchRequest) (*pb.Entry, error) {
	e, err := r.usecase.Fetch(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return r.convert(e), nil
}

func (r *OnGRPC) convert(e *domain.Entry) *pb.Entry {
	return &pb.Entry{
		Id: e.ID, Title: e.Title,
	}
}

func (r *OnGRPC) logf(format string, as ...interface{}) {
	logf(fmt.Sprintf("grpc: %s", format), as...)
}

type Printer interface {
	PrintIndex(io.Writer, []domain.Entry)
	Print(io.Writer, *domain.Entry)
}

func logf(format string, as ...interface{}) {
	if Logger == nil {
		return
	}

	Logger.Printf(format, as...)
}

var Logger logger

type logger interface {
	Printf(string, ...interface{})
}
