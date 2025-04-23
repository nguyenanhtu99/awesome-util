package mongoxv2

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReplaceOne executes an update command to replace at most one document in the collection.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/update/.
func ReplaceOne(
	ctx context.Context,
	collName string,
	filter interface{},
	model any,
	opts ...*options.ReplaceOptions,
) (*mongo.UpdateResult, error) {
	return CollWrite(collName).ReplaceOne(ctx, filter, model, opts...)
}

// ReplaceOneT executes an update command to replace at most one document in the collection.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/update/.
func ReplaceOneT[T Modeler](
	ctx context.Context,
	filter interface{},
	model T,
	opts ...*options.ReplaceOptions,
) (*mongo.UpdateResult, error) {
	return ReplaceOne(ctx, model.CollName(), filter, model, opts...)
}
