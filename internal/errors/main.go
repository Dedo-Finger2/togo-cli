package errors

import (
	"log/slog"
)

func ResourceCannotBeEmpty(resourceName string) {
	slog.Warn("Resource cannot be empty.", "resource", resourceName)
}

func ResourceNotFound(resourceName string) {
	slog.Error("Resource not found.", "resource", resourceName)
}
