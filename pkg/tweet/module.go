package tweet

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewRepository,
	NewService,
)
