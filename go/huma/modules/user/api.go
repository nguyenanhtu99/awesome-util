package user

import (
	"context"
	"huma/x"

	"github.com/danielgtaylor/huma/v2"
)

type UserHandler struct {
}

func New() *UserHandler {
	return &UserHandler{}
}

func (s *UserHandler) RegisterGetUser(api huma.API) {
	x.Register(api, huma.Operation{
		Method:      "GET",
		Path:        "/users/{userId}",
		Summary:     "Get a user",
		Description: "Get a user by ID",
		Tags:        []string{"User"},
	}, s.GetUser)
}

type GetUserInput struct {
	Body struct {
		Username string `json:"username" example:"Tuna" doc:"username"`
	}
	UserID string `path:"userId" example:"1" doc:"ID of the user"`
	Sort   string `query:"sort" example:"asc" doc:"Sort order"`
}

type GetUserOutput struct {
	Username string `json:"username" example:"Tuna" doc:"username"`
}

func (s *UserHandler) GetUser(ctx context.Context, input *GetUserInput) (*GetUserOutput, error) {
	
	resp := GetUserOutput{
		Username: "Tuna",
	}

	return &resp, nil
}
