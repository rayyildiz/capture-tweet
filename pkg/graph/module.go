package graph

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewResolver,
)
