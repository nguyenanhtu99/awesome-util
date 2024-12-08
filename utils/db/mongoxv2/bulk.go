package mongoxv2

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// BulkWrite performs a bulk write operation to the collection associated with the given collection name.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/bulkWrite/.
func BulkWrite(ctx context.Context, collName string, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	coll := CollWrite(collName)

	// Execute the bulk write command.
	return coll.BulkWrite(ctx, models, opts...)
}

// BulkWriteT performs a bulk write operation to the collection associated with the type T.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/bulkWrite/.
func BulkWriteT[T Modeler](ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	var t T
	return BulkWrite(ctx, t.CollName(), models, opts...)
}
