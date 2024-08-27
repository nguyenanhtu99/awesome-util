package mongox

import (
	"context"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo CommonRepo[T]) Update(ctx context.Context, data T, opts ...*options.UpdateOptions) error {
	model := repo.makeEmptyModel()
	coll := mgm.Coll(model)

	err := coll.UpdateWithCtx(ctx, data, opts...)
	if err != nil {
		return err
	}

	return nil
}
func (repo CommonRepo[T]) UpdateOneById(ctx context.Context, id string, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: objId}}

	return repo.UpdateOne(ctx, filter, update, opts...)
}

func (repo CommonRepo[T]) UpdateOne(ctx context.Context, filter, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	model := repo.makeEmptyModel()
	coll := mgm.Coll(model)

	return coll.UpdateOne(ctx, filter, update, opts...)
}

func (repo CommonRepo[T]) UpdateMany(ctx context.Context, filter, update any, opts ...*options.UpdateOptions) (int64, error) {
	model := repo.makeEmptyModel()
	coll := mgm.Coll(model)

	modifyDocument := int64(0)
	result, err := coll.UpdateMany(ctx, filter, update, opts...)
	if err != nil {
		return modifyDocument, err
	}
	modifyDocument = result.ModifiedCount

	return modifyDocument, nil
}

func (repo CommonRepo[T]) FindOneAndUpdate(ctx context.Context, filter interface{},
	update interface{}, opts ...*options.FindOneAndUpdateOptions) (T, error) {

	model := repo.makeEmptyModel()
	coll := mgm.Coll(model)

	res := coll.FindOneAndUpdate(ctx, filter, update, opts...)
	err := res.Decode(model)
	if err != nil {
		return model, err
	}

	return model, nil
}
