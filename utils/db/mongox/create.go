package mongox

import (
	"context"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo CommonRepo[T]) Create(ctx context.Context, data T, opts ...*options.InsertOneOptions) error {
	model := repo.makeEmptyModel()
	coll := mgm.Coll(model)

	err := coll.CreateWithCtx(ctx, data, opts...)
	if err != nil {
		return err
	}

	return nil
}

func (repo CommonRepo[T]) CreateMany(ctx context.Context, data []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	model := repo.makeEmptyModel()
	coll := mgm.Coll(model)

	return coll.InsertMany(ctx, data, opts...)
}
