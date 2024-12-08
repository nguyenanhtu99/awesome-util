package mongoxv2

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// CountDocuments returns the number of documents in the collection.
// For a fast count of the documents in the collection, see the EstimatedCount function.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/method/db.collection.countDocuments/.
func Count(ctx context.Context, collName string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	// Retrieve the collection for reading using the collection name.
	collection := CollRead(collName)

	// Count and return the number of documents that match the filter.
	return collection.CountDocuments(ctx, filter, opts...)
}

// CountT counts the number of documents in the collection associated with the type T that match the given filter. 
// For a fast count of the documents in the collection, see the EstimatedCountT function.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/method/db.collection.countDocuments/.
func CountT[T Modeler](ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	var t T
	return Count(ctx, t.CollName(), filter, opts...)
}

// EstimatedCount returns an estimate of the number of documents in the collection
// associated with the given collection name using collection metadata.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/method/db.collection.estimatedDocumentCount/.
func EstimatedCount(ctx context.Context, collName string, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	// Retrieve the collection for reading using the collection name.
	collection := CollRead(collName)

	// Execute the estimated document count command and return the result.
	return collection.EstimatedDocumentCount(ctx, opts...)
}

// EstimatedCountT returns an estimate of the number of documents in the collection associated with the type T.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/method/db.collection.estimatedDocumentCount/.
func EstimatedCountT[T Modeler](ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	var t T
	// Retrieve the collection for reading using the collection name associated with the type T.
	return EstimatedCount(ctx, t.CollName(), opts...)
}
