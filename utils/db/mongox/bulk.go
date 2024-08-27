package mongox

import (
	"context"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo CommonRepo[T]) BulkWrite(ctx context.Context, models []mongo.WriteModel,
	opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	model := repo.makeEmptyModel()
	coll := mgm.Coll(model)

	return coll.BulkWrite(ctx, models, opts...)
}
