package userrepo

import (
	"awesome-util/graphql/gqlgen/modules/user/usermodel"
	"awesome-util/utils/db/mongox"
	"context"
)

func New() *Repository {
	return &Repository{
		mongo: mongox.NewCommonRepo[*usermodel.User](),
	}
}

type Repository struct {
	mongo mongox.BaseRepository[*usermodel.User]
}

func (r *Repository) Create(ctx context.Context, user *usermodel.User) error {
	return r.mongo.Create(ctx, user)
}

func (r *Repository) Get(ctx context.Context, id string) (*usermodel.User, error) {
	return r.mongo.FindById(ctx, id)
}
