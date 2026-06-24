package main

import (
	"bytes"
	"log/slog"
)

func Test_requestLogger(t *testing.T) {
	logBuffer := &bytes.Buffer{}

	logger := slog.New(slog.NewTextHandler(logBuffer, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr) slog,
	}))
}
