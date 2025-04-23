package mongoxv2

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AddIndex adds an index to the specified collection.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/method/db.collection.createIndex/.
func AddIndex(
	ctx context.Context,
	collName string,
	index mongo.IndexModel,
	opts ...*options.CreateIndexesOptions,
) (string, error) {
	return CollWrite(collName).Indexes().CreateOne(ctx, index, opts...)
}

// AddIndexT adds an index to the collection associated with the type T.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/method/db.collection.createIndex/.
func AddIndexT[T Modeler](
	ctx context.Context,
	index mongo.IndexModel,
	opts ...*options.CreateIndexesOptions,
) (string, error) {
	var t T

	return AddIndex(ctx, t.CollName(), index, opts...)
}

// AddIndexes adds multiple indexes to the collection associated with the given collection name.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/method/db.collection.createIndexes/.
func AddIndexes(
	ctx context.Context,
	collName string,
	indexes []mongo.IndexModel,
	opts ...*options.CreateIndexesOptions,
) ([]string, error) {
	return CollWrite(collName).Indexes().CreateMany(ctx, indexes, opts...)
}

// AddIndexesT adds multiple indexes to the collection associated with the type T.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/method/db.collection.createIndexes/.
func AddIndexesT[T Modeler](
	ctx context.Context,
	indexes []mongo.IndexModel,
	opts ...*options.CreateIndexesOptions,
) ([]string, error) {
	var t T

	return AddIndexes(ctx, t.CollName(), indexes, opts...)
}
