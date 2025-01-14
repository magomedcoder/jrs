package cli

import (
	"github.com/google/wire"
	"github.com/magomedcoder/gskeleton/internal/cli/handler"
	"github.com/magomedcoder/gskeleton/internal/infrastructure"
	"github.com/magomedcoder/gskeleton/internal/usecase"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),
	wire.Struct(new(handler.Migrate), "*"),
	usecase.ProviderSet,
	infrastructure.ProviderSet,
)
