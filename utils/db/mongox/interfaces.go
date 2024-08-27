package mongox

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo/readconcern"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type BaseRepository[T mgm.Model] interface {
	FindById(ctx context.Context, id string) (T, error)
	FindByListId(ctx context.Context, arrId []string, opts ...*options.FindOptions) ([]T, error)
	Find(ctx context.Context, filter any, opts ...*options.FindOneOptions) (T, error)
	FindMany(ctx context.Context, filter any, opts ...*options.FindOptions) ([]T, error)
	FindOneAndUpdate(ctx context.Context, filter interface{},
		update interface{}, opts ...*options.FindOneAndUpdateOptions) (T, error)
	Create(ctx context.Context, data T, opts ...*options.InsertOneOptions) error
	CreateMany(ctx context.Context, data []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	Update(ctx context.Context, data T, opts ...*options.UpdateOptions) error
	UpdateOne(ctx context.Context, filter, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateOneById(ctx context.Context, id string, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter, update any, opts ...*options.UpdateOptions) (int64, error)
	FindRandom(ctx context.Context, filter any, size int, opts ...*options.AggregateOptions) ([]bson.M, error)
	BulkWrite(ctx context.Context, models []mongo.WriteModel,
		opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error)
	Count(ctx context.Context, filter any, opts ...*options.CountOptions) (int64, error)
	EstimatedCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error)
	Aggregate(ctx context.Context, results any, stages []any, opts ...*options.AggregateOptions) error
	FindOneAndReplace(ctx context.Context, filter any, replacement any, opts ...*options.FindOneAndReplaceOptions) (T, error)
}

type CommonRepo[T mgm.Model] struct{}

func NewCommonRepo[T mgm.Model]() CommonRepo[T] {
	return CommonRepo[T]{}
}

// Validate interface implement at compile time
var _ BaseRepository[mgm.Model] = (*CommonRepo[mgm.Model])(nil)

func (repo CommonRepo[T]) makeEmptyModel() T {
	var model T

	typeM := reflect.TypeOf(model)
	if typeM.Kind() == reflect.Ptr {
		return reflect.New(typeM.Elem()).Interface().(T)
	}
	return model
}

func CollRead(model mgm.Model, opts ...*options.CollectionOptions) *mgm.Collection {
	opt := options.Collection()
	opt.SetReadPreference(readpref.Nearest())
	opt.SetReadConcern(readconcern.Majority())
	opts = append(opts, opt)

	coll := mgm.Coll(model, opts...)
	return coll
}

type Indexer interface {
	Index(ctx context.Context) error
}

func CreateIndex(ctx context.Context, handleErr func(err error), repos ...interface{}) {
	for i := range repos {
		if repo, ok := repos[i].(Indexer); ok {
			err := repo.Index(ctx)
			if err != nil {
				handleErr(err)
			}
		}
	}
}
