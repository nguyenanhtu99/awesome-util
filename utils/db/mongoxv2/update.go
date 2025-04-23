package mongoxv2

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateOne executes an update command to update at most one document in the collection.
//
// The opts parameter can be used to specify options for the operation (see the options.UpdateOptions documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/update/.
func UpdateOne(
	ctx context.Context,
	collName string,
	filter interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return CollWrite(collName).UpdateOne(ctx, filter, update, opts...)
}

// UpdateOneT executes an update command to update at most one document in the collection.
//
// The opts parameter can be used to specify options for the operation (see the options.UpdateOptions documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/update/.
func UpdateOneT[T Modeler](
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	var t T
	return UpdateOne(ctx, t.CollName(), filter, update, opts...)
}

// UpdateAndReturn executes an update command to update at most one document in the collection and returns the document.
//
// When ReturnDocument in the options is set to After, the function returns the updated document.
// The default is Before, which means the original document will be returned before the replacement is performed.
// The opts parameter can be used to specify options for the operation (see the options.FindOneAndUpdateOptions documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/findAndModify/.
func UpdateAndReturn[T Modeler](
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.FindOneAndUpdateOptions,
) (T, error) {
	var t T
	res := CollWrite(t.CollName()).FindOneAndUpdate(ctx, filter, update, opts...)

	err := res.Decode(&t)
	return t, err
}

// UpdateMany executes an update command to update one or more documents in the collection.
//
// The opts parameter can be used to specify options for the operation (see the options.UpdateOptions documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/update/.
func UpdateMany(
	ctx context.Context,
	collName string,
	filter interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return CollWrite(collName).UpdateMany(ctx, filter, update, opts...)
}

// UpdateManyT executes an update command to update one or more documents in the collection.
//
// The opts parameter can be used to specify options for the operation (see the options.UpdateOptions documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/update/.
func UpdateManyT[T Modeler](
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	var t T
	return UpdateMany(ctx, t.CollName(), filter, update, opts...)
}
