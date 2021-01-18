package infra

import (
	"os"
)

func IsDebug() bool {
	debug := os.Getenv("DEBUG")
	if debug == "true" || debug == "TRUE" {
		return true
	}

	return false
}
