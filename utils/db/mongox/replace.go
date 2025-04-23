package mongox

import (
	"context"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo CommonRepo[T]) FindOneAndReplace(
	ctx context.Context,
	filter any,
	replacement any,
	opts ...*options.FindOneAndReplaceOptions,
) (T, error) {
	model := repo.makeEmptyModel()
	coll := mgm.Coll(model)

	res := coll.FindOneAndReplace(ctx, filter, replacement, opts...)
	err := res.Decode(model)
	if err != nil {
		return model, err
	}

	return model, nil
}
