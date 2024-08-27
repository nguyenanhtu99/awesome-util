package main

import (
	"awesome-util/graphql/go-graphql/modules/user/userapi"
	"awesome-util/utils/db/mongox"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Connect to MongoDB
	initCtx := context.Background()
	if err := mongox.ConnectMongo(initCtx); err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if err := initHTTP(e.Group("/v1")); err != nil {
		fmt.Printf("failed to init http: %v", err)
	}

	// GraphQL handler
	e.POST("/graphql", func(c echo.Context) error {
		var params struct {
			Query string `json:"query"`
		}
		if err := c.Bind(&params); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: params.Query,
		})

		if len(result.Errors) > 0 {
			return c.JSON(http.StatusBadRequest, result.Errors)
		}

		return c.JSON(http.StatusOK, result)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func initHTTP(g *echo.Group) error {
	userGroup := g.Group("/users")
	if err := userapi.New(userGroup); err != nil {
		return fmt.Errorf("failed to init user api: %w", err)
	}

	return nil
}
