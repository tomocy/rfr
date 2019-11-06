package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tomocy/rfv/cmd/rfv/format"
	"github.com/tomocy/rfv/cmd/rfv/mode"
)

func main() {
	cnf := parseConfig()
	if err := cnf.newRunner().Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func parseConfig() *config {
	m, f := flag.String("m", "http", "mode"), flag.String("f", "json", "format")
	addr := flag.String("addr", ":80", "address")
	log := flag.String("log", "", "log level")
	flag.Parse()

	return &config{
		mode: *m, format: *f,
		addr: *addr,
		log:  *log,
	}
}

type config struct {
	mode, format string
	addr         string
	log          string
}

func (c *config) newRunner() runner {
	c.setLogger()

	switch c.mode {
	case modeHTTP:
		return mode.NewOnHTTP(c.addr, c.newPrinter())
	case modeGRPC:
		return mode.NewOnGRPC(c.addr)
	default:
		return new(help)
	}
}

func (c *config) setLogger() {
	switch c.log {
	case logDebug:
		mode.Logger = log.New(os.Stdout, "", log.LstdFlags)
	}
}

const logDebug = "debug"

const (
	modeHTTP = "http"
	modeGRPC = "grpc"
)

func (c *config) newPrinter() mode.Printer {
	switch c.format {
	case formatJSON:
		return new(format.InJSON)
	default:
		return nil
	}
}

const formatJSON = "json"

type runner interface {
	Run() error
}

type help struct {
	err error
}

func (h *help) Run() error {
	flag.Usage()

	return h.err
}
