//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/google/wire/cmd/wire"
)
