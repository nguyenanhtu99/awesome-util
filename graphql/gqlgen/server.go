package main

import (
	"awesome-util/graphql/gqlgen/graph"
	"awesome-util/graphql/gqlgen/modules/user/userapi"
	"awesome-util/utils/db/mongox"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const defaultPort = "8080"

func main() {
	// Connect to MongoDB
	initCtx := context.Background()
	if err := mongox.ConnectMongo(initCtx); err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	v1 := e.Group("/v1")

	resolver, err := graph.New()
	if err != nil {
		log.Fatalf("failed to init resolver: %v", err)
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	v1.POST("/graphql", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	v1.GET("/graphql", func(c echo.Context) error {
		playground.Handler("GraphQL playground", "/graphql").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	if err := initHTTP(v1); err != nil {
		log.Fatalf("failed to init http: %v", err)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(e.Start(":" + port))
}

func initHTTP(g *echo.Group) error {
	userGroup := g.Group("/users")
	if err := userapi.New(userGroup); err != nil {
		return fmt.Errorf("failed to init user api: %w", err)
	}

	return nil
}
