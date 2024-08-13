package utils

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/Dedo-Finger2/todo-list-cli/internal/errors"
)

// "taskID", ["not-null", "integer"]
// "toGoListName", ["not-null"]

func Validator(name string, value *string, rules []string) {
	for idx := 0; idx < len(rules); idx++ {
		switch {
		case rules[idx] == "not-null":
			validateEmptyValues(name, value)
		case rules[idx] == "integer":
			validateStringCanBeConvertedIntoInteger(value)
		default:
		}
	}
}

func validateEmptyValues(name string, value *string) {
	if *value == "" {
		errors.ResourceCannotBeEmpty(name)
		os.Exit(1)
	}
}

func validateStringCanBeConvertedIntoInteger(value *string) {
	if convertedValue, err := strconv.Atoi(*value); err != nil || convertedValue == 0 {
		slog.Error("Error trying to convert string into integer.", "error", err)
		os.Exit(1)
	}
}
