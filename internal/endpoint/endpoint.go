package endpoint

import (
	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/fx"
)

type Endpoint interface {
	RegisterRoute(openapi huma.API)
}

func Annotation[T any](a T) interface{} {
	return fx.Annotate(
		a,
		fx.As(new(Endpoint)),
		fx.ResultTags(`group:"endpoint"`),
	)
}
