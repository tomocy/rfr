package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	cnf := parseConfig()
	if err := cnf.newRunner().Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func parseConfig() *config {
	return new(config)
}

type config struct {
	mode, format string
}

func (c *config) newRunner() runner {
	return new(help)
}

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
