package main

import (
	"context"
	"fmt"
	"huma/modules/user"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/labstack/echo/v4"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

// GreetingInput represents the greeting operation request.
type GreetingInput struct {
	Name string `path:"name" example:"Tuna" doc:"Name of the person"`
}

func main() {
	// Create a new router & API.
	e := echo.New()
	config := huma.DefaultConfig("huma", "1.0.0")
	api := humaecho.New(e, config)

	huma.Get(api, "/greeting/{name}", func(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})

	userHandler := user.New()
	huma.AutoRegister(api, userHandler)

	e.Start("localhost:8080")
}
