package search

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewService,
)
