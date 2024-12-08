package mongoxv2

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindManyT executes a query to find documents in the collection associated with the type T.
//
// The opts parameter can be used to specify options for the operation (see the options.FindOptions documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/find/.
func FindManyT[T Modeler](ctx context.Context, filter interface{}, opts ...*options.FindOptions) (res []T, err error) {
	var t T
	// Create a cursor from the collection associated with the type T
	cur, err := CollRead(t.CollName()).Find(ctx, filter, opts...)
	if err != nil {
		return res, err
	}

	// Decode all the documents in the cursor into the slice of T
	err = cur.All(ctx, &res)
	return res, err
}

// FindOneT executes a query to find one document in the collection associated with the type T.
//
// The opts parameter can be used to specify options for the operation (see the options.FindOneOptions documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/find/.
func FindOneT[T Modeler](ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (T, error) {
	var t T

	// Find one document in the collection with the given filter
	cur := CollRead(t.CollName()).FindOne(ctx, filter, opts...)

	// Decode the document into the given pointer
	err := cur.Decode(&t)

	return t, err
}
