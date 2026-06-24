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

		return slog.New(slog.NewTextHandler(multiWriter,nil)), close, nil
	}

	close := func() error {
		return nil
	}

	return slog.New(slog.NewTextHandler(os.Stderr, nil)), close ,nil
}
