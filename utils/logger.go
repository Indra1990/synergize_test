package utils

import (
	"log/slog"

	"github.com/go-chi/httplog/v2"
)

func LoggerChi() *httplog.Logger{
	logger := httplog.NewLogger("app-synergize",httplog.Options{
		LogLevel: slog.LevelDebug,
		Concise: true,
		RequestHeaders:   true,
		MessageFieldName: "message",

		// TimeFieldFormat: time.RFC850,
		Tags: map[string]string{
			"version": "v1.0-81aa4244d9fc8076a",
			"env":     "dev",
		  },

	})

	return logger
}