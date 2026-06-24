package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
)

type closeFunc func() error

func initializeLogger(logFile string) (*slog.Logger, closeFunc, error) {
	debugHandler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open log file: %w", err)
		}
		bufferedFile := bufio.NewWriterSize(file, 8192)
		multiWriter := io.MultiWriter(os.Stderr, bufferedFile)

		close := func() error {
			if err := bufferedFile.Flush(); err != nil {
				return fmt.Errorf("failed to flush log file: %w", err)
			}

			if err := file.Close(); err != nil {
				fmt.Errorf("failed to close log file: %w", err)

			}
			return nil
		}

		infoHandler := slog.NewJSONHandler(multiWriter, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})

		logger := slog.New(slog.NewMultiHandler(
			debugHandler,
			infoHandler,
			))

		return logger, close, nil
	}

	infoHandler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	logger := slog.New(slog.NewMultiHandler(
		debugHandler,
		infoHandler,
		))

	close := func() error {
		return nil
	}

	return logger, close ,nil
}
