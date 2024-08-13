package errors

import (
	"log/slog"
	"os"
)

func ResourceCannotBeEmpty(resourceName string) {
	slog.Warn("Resource cannot be empty.", "resource", resourceName)
	os.Exit(1)
}

func ResourceNotFound(resourceName string) {
	slog.Error("Resource not found.", "resource", resourceName)
	os.Exit(1)
}
