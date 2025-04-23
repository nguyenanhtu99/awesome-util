package mongoxv2

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create inserts a single document into the specified collection.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/method/db.collection.insertOne/.
func Create(ctx context.Context, collName string, model any, opts ...*options.InsertOneOptions) (interface{}, error) {
	col := CollWrite(collName)

	result, err := col.InsertOne(ctx, model, opts...)
	if err != nil {
		return nil, err
	}

	// Return the inserted ID to the caller.
	return result.InsertedID, nil
}

// CreateT inserts a single document into the collection associated with the model. The function returns the inserted ID to the caller.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/method/db.collection.insertOne/.
func CreateT[T Modeler](ctx context.Context, model T, opts ...*options.InsertOneOptions) (interface{}, error) {
	return Create(ctx, model.CollName(), model, opts...)
}

// CreateWithID inserts a single document into the collection associated with the model.
// It sets the ID field of the model with the inserted document's ID.
func CreateWithID[T IDModeler](ctx context.Context, model T, opts ...*options.InsertOneOptions) error {
	// Get the collection for writing using the model's collection name.
	col := CollWrite(model.CollName())

	// Insert the document into the collection.
	result, err := col.InsertOne(ctx, model, opts...)
	if err != nil {
		// Return the error if the insert operation fails.
		return err
	}

	// Set the ID field of the model with the inserted document's ID.
	model.SetID(result.InsertedID)
	return nil
}

// CreateMany inserts multiple documents into the collection associated with the collName.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/method/db.collection.insertMany/.
func CreateMany(
	ctx context.Context,
	collName string,
	models []any,
	opts ...*options.InsertManyOptions,
) (*mongo.InsertManyResult, error) {
	if len(models) == 0 {
		return &mongo.InsertManyResult{}, nil
	}

	// Get the collection for writing using the model's collection name.
	col := CollWrite(collName)

	// Insert the documents into the collection.
	return col.InsertMany(ctx, models, opts...)
}

// CreateManyT inserts multiple documents into the collection associated with the type T.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/method/db.collection.insertMany/.
func CreateManyT[T Modeler](
	ctx context.Context,
	models []T,
	opts ...*options.InsertManyOptions,
) (*mongo.InsertManyResult, error) {
	if len(models) == 0 {
		return &mongo.InsertManyResult{}, nil
	}

	// Convert the slice of T to a slice of interface{}.
	// This is necessary because the InsertMany method takes a slice of interface{}
	// as an argument.
	m := make([]interface{}, len(models))
	for i := range models {
		m[i] = models[i]
	}

	// Call the CreateMany function with the slice of interface{}
	// and the collection name associated with the type T.
	return CreateMany(ctx, models[0].CollName(), m, opts...)
}
