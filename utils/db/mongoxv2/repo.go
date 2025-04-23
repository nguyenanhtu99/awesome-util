package mongoxv2

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repoer[T IDModeler] interface {
	CreateOne(ctx context.Context, data T, opts ...*options.InsertOneOptions) error
	CreateMany(ctx context.Context, data []T, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)

	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (T, error)
	FindMany(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]T, error)

	UpdateOne(
		ctx context.Context,
		filter interface{},
		update interface{},
		opts ...*options.UpdateOptions,
	) (*mongo.UpdateResult, error)
	UpdateMany(
		ctx context.Context,
		filter interface{},
		update interface{},
		opts ...*options.UpdateOptions,
	) (*mongo.UpdateResult, error)
	UpdateAndReturn(
		ctx context.Context,
		filter interface{},
		update interface{},
		opts ...*options.FindOneAndUpdateOptions,
	) (T, error)
	ReplaceOne(
		ctx context.Context,
		filter interface{},
		data T,
		opts ...*options.ReplaceOptions,
	) (*mongo.UpdateResult, error)

	Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
	EstimatedCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error)

	Aggregate(ctx context.Context, results interface{}, pipeline []interface{}, opts ...*options.AggregateOptions) error
	BulkWrite(
		ctx context.Context,
		models []mongo.WriteModel,
		opts ...*options.BulkWriteOptions,
	) (*mongo.BulkWriteResult, error)
}

type Repo[T IDModeler] struct{}

func NewRepo[T IDModeler]() Repo[T] {
	return Repo[T]{}
}

var _ Repoer[IDModeler] = (*Repo[IDModeler])(nil)

func (repo Repo[T]) CreateOne(ctx context.Context, data T, opts ...*options.InsertOneOptions) error {
	return CreateWithID(ctx, data, opts...)
}

func (repo Repo[T]) CreateMany(
	ctx context.Context,
	data []T,
	opts ...*options.InsertManyOptions,
) (*mongo.InsertManyResult, error) {
	return CreateManyT(ctx, data, opts...)
}

func (repo Repo[T]) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (T, error) {
	return FindOneT[T](ctx, filter, opts...)
}

func (repo Repo[T]) FindMany(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]T, error) {
	return FindManyT[T](ctx, filter, opts...)
}

func (repo Repo[T]) UpdateOne(
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return UpdateOneT[T](ctx, filter, update, opts...)
}

func (repo Repo[T]) UpdateMany(
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return UpdateManyT[T](ctx, filter, update, opts...)
}

func (repo Repo[T]) UpdateAndReturn(
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.FindOneAndUpdateOptions,
) (T, error) {
	return UpdateAndReturn[T](ctx, filter, update, opts...)
}

func (repo Repo[T]) ReplaceOne(
	ctx context.Context,
	filter interface{},
	data T,
	opts ...*options.ReplaceOptions,
) (*mongo.UpdateResult, error) {
	return ReplaceOneT(ctx, filter, data, opts...)
}

func (repo Repo[T]) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return CountT[T](ctx, filter, opts...)
}

func (repo Repo[T]) EstimatedCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return EstimatedCountT[T](ctx, opts...)
}

func (repo Repo[T]) Aggregate(
	ctx context.Context,
	results interface{},
	pipeline []interface{},
	opts ...*options.AggregateOptions,
) error {
	return AggregateT[T](ctx, results, pipeline, opts...)
}

func (repo Repo[T]) BulkWrite(
	ctx context.Context,
	models []mongo.WriteModel,
	opts ...*options.BulkWriteOptions,
) (*mongo.BulkWriteResult, error) {
	return BulkWriteT[T](ctx, models, opts...)
}
