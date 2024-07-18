package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/stmcginnis/apiserver/internal/api"
)

const (
	httpListenAddress = ":8080"
	logLevel          = slog.LevelDebug
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel,
	}))
	slog.SetDefault(logger)

	mux := http.DefaultServeMux
	for _, endpoint := range api.RegisteredEndpoints() {
		slog.Info("serving endpoint", "endpoint", endpoint.String())
	}

	err := http.ListenAndServe(httpListenAddress, mux)
	if err != nil {
		slog.Error("error running http server", "err", err)
	}
}
