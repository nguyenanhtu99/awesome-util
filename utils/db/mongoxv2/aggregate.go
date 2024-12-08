package mongoxv2

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Aggregate executes an aggregate command against the collection and returns a cursor over the resulting documents.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/aggregate/.
func Aggregate(ctx context.Context, collName string, pipeline []interface{}, opts ...*options.AggregateOptions) (res []bson.M, err error) {
	// Create a cursor from the collection associated with the type T
	col := db.Collection(collName)

	// Execute the aggregate command
	cursor, err := col.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}

	// Decode all the documents in the cursor into the slice of bson.M
	err = cursor.All(ctx, &res)

	return res, err
}

// AggregateT executes an aggregate command against the collection associated with the type T and returns a cursor over the
// resulting documents.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/aggregate/.
func AggregateT[T Modeler](ctx context.Context, pipeline []interface{}, opts ...*options.AggregateOptions) (res []bson.M, err error) {
	// Create a model value of type T
	var t T

	// Aggregate the documents in the collection associated with the type T
	return Aggregate(ctx, t.CollName(), pipeline, opts...)
}
