package mongoxv2

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// Aggregate executes an aggregate command against the collection and returns a cursor over the resulting documents.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/aggregate/.
func Aggregate(ctx context.Context, collName string, result interface{}, pipeline []interface{}, opts ...*options.AggregateOptions) error {
	// Create a cursor from the collection associated with the type T
	col := db.Collection(collName)

	// Execute the aggregate command
	cursor, err := col.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return err
	}

	// Decode all the documents in the cursor into the result
	return cursor.All(ctx, result)
}

// AggregateT executes an aggregate command against the collection associated with the type T and returns a cursor over the
// resulting documents.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/aggregate/.
func AggregateT[T Modeler](ctx context.Context, result interface{}, pipeline []interface{}, opts ...*options.AggregateOptions) error {
	var t T

	return Aggregate(ctx, t.CollName(), result, pipeline, opts...)
}
