package mode

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"

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
		usecase: *app.NewRFCUsecase(new(infra.ViaHTTP)),
		printer: printer,
	}
}

type OnHTTP struct {
	addr    string
	router  chi.Router
	usecase app.RFCUsecase
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
	r.router.Get("/", r.get)
	r.router.Get("/{id}", r.find)
}

func (r *OnHTTP) get(w http.ResponseWriter, _ *http.Request) {
	got, err := r.usecase.Get(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.printer.PrintAll(w, got); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (r *OnHTTP) find(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	found, err := r.usecase.Find(context.Background(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.printer.Print(w, found); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (r *OnHTTP) logf(format string, as ...interface{}) {
	logf(fmt.Sprintf("http: %s", format), as...)
}

func NewOnGRPC(addr string) *OnGRPC {
	return &OnGRPC{
		addr:    addr,
		server:  *grpc.NewServer(),
		usecase: *app.NewRFCUsecase(new(infra.ViaHTTP)),
	}
}

type OnGRPC struct {
	pb.UnimplementedRFCRepoServer
	addr    string
	server  grpc.Server
	usecase app.RFCUsecase
}

func (r *OnGRPC) Run() error {
	pb.RegisterRFCRepoServer(&r.server, r)

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

func (r *OnGRPC) Get(ctx context.Context, req *empty.Empty) (*pb.RFCs, error) {
	got, err := r.usecase.Get(ctx)
	if err != nil {
		return nil, err
	}

	converteds := make([]*pb.RFC, len(got))
	for i, rfc := range got {
		converteds[i] = r.convert(rfc)
	}

	return &pb.RFCs{
		Rfcs: converteds,
	}, nil
}

func (r *OnGRPC) Find(ctx context.Context, req *pb.FindRequest) (*pb.RFC, error) {
	found, err := r.usecase.Find(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return r.convert(found), nil
}

func (r *OnGRPC) convert(rfc *domain.RFC) *pb.RFC {
	return &pb.RFC{
		Id:    int32(rfc.ID),
		Title: rfc.Title,
	}
}

func (r *OnGRPC) logf(format string, as ...interface{}) {
	logf(fmt.Sprintf("grpc: %s", format), as...)
}

type Printer interface {
	PrintAll(io.Writer, []*domain.RFC) error
	Print(io.Writer, *domain.RFC) error
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
