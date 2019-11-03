package main

import (
	"flag"
	"fmt"
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
	flag.Parse()

	return &config{
		mode: *m, format: *f,
	}
}

type config struct {
	mode, format string
}

func (c *config) newRunner() runner {
	switch c.mode {
	case modeHTTP:
		return mode.NewOnHTTP(":80", c.newPrinter())
	default:
		return new(help)
	}
}

const modeHTTP = "http"

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
