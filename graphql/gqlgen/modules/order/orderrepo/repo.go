package orderrepo

import (
	"awesome-util/graphql/gqlgen/modules/order/ordermodel"
	"awesome-util/utils/db/mongox"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func New() *Repository {
	return &Repository{
		mongo: mongox.NewCommonRepo[*ordermodel.Order](),
	}
}

type Repository struct {
	mongo mongox.BaseRepository[*ordermodel.Order]
}

func (r *Repository) Create(ctx context.Context, user *ordermodel.Order) error {
	return r.mongo.Create(ctx, user)
}

func (r *Repository) Get(ctx context.Context, id string) (*ordermodel.Order, error) {
	return r.mongo.FindById(ctx, id)
}

func (r *Repository) GetByUserId(ctx context.Context, userId string) ([]*ordermodel.Order, error) {
	return r.mongo.FindMany(ctx, bson.M{"userId": userId})
}
