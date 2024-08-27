package mongox

import (
	"context"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/builder"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo CommonRepo[T]) FindById(ctx context.Context, id string) (T, error) {
	model := repo.makeEmptyModel()
	coll := CollRead(model)

	err := coll.FindByIDWithCtx(ctx, id, model)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (repo CommonRepo[T]) FindByListId(ctx context.Context, arrId []string, opts ...*options.FindOptions) ([]T, error) {
	oids := make([]primitive.ObjectID, 0)

	for _, id := range arrId {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			continue
		}
		oids = append(oids, objId)
	}

	filter := bson.M{"_id": bson.M{"$in": oids}}
	return repo.FindMany(ctx, filter, opts...)
}

func (repo CommonRepo[T]) Find(ctx context.Context, filter any, opts ...*options.FindOneOptions) (T, error) {
	model := repo.makeEmptyModel()
	coll := CollRead(model)

	err := coll.FirstWithCtx(ctx, filter, model, opts...)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (repo CommonRepo[T]) FindWithCollOption(ctx context.Context, filter any, findOpts []*options.FindOneOptions, collOpts []*options.CollectionOptions) (T, error) {
	model := repo.makeEmptyModel()
	coll := mgm.Coll(model, collOpts...)

	err := coll.FirstWithCtx(ctx, filter, model, findOpts...)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (repo CommonRepo[T]) FindMany(ctx context.Context, filter any, opts ...*options.FindOptions) ([]T, error) {
	model := repo.makeEmptyModel()
	coll := CollRead(model)

	results := make([]T, 0)
	err := coll.SimpleFindWithCtx(ctx, &results, filter, opts...)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (repo CommonRepo[T]) FindManyWithCollOption(ctx context.Context, filter any, findOpts []*options.FindOptions, collOpts []*options.CollectionOptions) ([]T, error) {
	model := repo.makeEmptyModel()
	coll := mgm.Coll(model, collOpts...)

	results := make([]T, 0)
	err := coll.SimpleFindWithCtx(ctx, &results, filter, findOpts...)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (repo CommonRepo[T]) FindRandom(ctx context.Context, filter any, size int, opts ...*options.AggregateOptions) ([]bson.M, error) {
	model := repo.makeEmptyModel()
	coll := mgm.Coll(model)
	pipeline := []bson.M{
		{
			"$match": filter,
		},
		{
			// Random items in database
			"$sample": bson.M{
				"size": size,
			},
		},
	}
	cursor, err := coll.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}
	var showsWithInfo []bson.M
	if err = cursor.All(ctx, &showsWithInfo); err != nil {
		return nil, err
	}
	return showsWithInfo, nil
}

func (repo CommonRepo[T]) Count(ctx context.Context, filter any, opts ...*options.CountOptions) (int64, error) {
	model := repo.makeEmptyModel()
	coll := CollRead(model)
	return coll.CountDocuments(ctx, filter, opts...)
}

func (repo CommonRepo[T]) EstimatedCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	model := repo.makeEmptyModel()
	coll := CollRead(model)
	return coll.EstimatedDocumentCount(ctx, opts...)
}

func (repo CommonRepo[T]) CountWithCollOption(ctx context.Context, filter any, countOpts []*options.CountOptions, collOpts []*options.CollectionOptions) (int64, error) {
	model := repo.makeEmptyModel()
	coll := mgm.Coll(model, collOpts...)
	return coll.CountDocuments(ctx, filter, countOpts...)
}

func (repo CommonRepo[T]) Aggregate(ctx context.Context, results any, stages []any, opts ...*options.AggregateOptions) error {
	model := repo.makeEmptyModel()
	coll := CollRead(model)

	pipeline := bson.A{}
	for _, stage := range stages {
		if operator, ok := stage.(builder.Operator); ok {
			pipeline = append(pipeline, builder.S(operator))
		} else {
			pipeline = append(pipeline, stage)
		}
	}

	cursor, err := coll.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return err
	}

	return cursor.All(ctx, results)
}
