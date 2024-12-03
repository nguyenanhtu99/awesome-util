package x

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
)

type Output struct {
	Body any `json:",inline" doc:",inline"`
}

func Register[I, O any](api huma.API, op huma.Operation, h func(context.Context, *I) (*O, error)) {
	huma.Register(api, op, func(ctx context.Context, i *I) (*Output, error) {
		output, err := h(ctx, i)

		return &Output{Body: output}, err
	})
}
