package utils

import (
	"flag"
	"strings"
)

func DefineFlagValue(flagName *string) {
	if strings.Contains(flag.Arg(1), "=") {
		*flagName = strings.Split(flag.Arg(1), "=")[1]
	} else {
		*flagName = flag.Arg(2)
	}
}
