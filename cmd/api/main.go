package main

import (
	"com.capturetweet/internal/infra"
	"com.capturetweet/pkg/graph"
	"com.capturetweet/pkg/graph/generated"
	"com.capturetweet/pkg/search"
	"com.capturetweet/pkg/tweet"
	"com.capturetweet/pkg/user"
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"net/http"
	"os"
)

func init() {
	godotenv.Load()
}

func Register(resolver generated.ResolverRoot) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	if os.Getenv("GRAPHQL_ENABLE_PLAYGROUND") == "true" {
		http.Handle("/", playground.Handler("GraphQL playground", "/api/query"))
	}
	http.Handle("/api/query", srv)
}

func InitServer(lifecycle fx.Lifecycle) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	srv := &http.Server{
		Addr: ":" + port,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Close()
		},
	})
}

func main() {
	app := fx.New(
		infra.Module,  // Database, Logger
		search.Module, // Search Module
		user.Module,   // user repository, service
		tweet.Module,  // tweet repository, service
		graph.Module,  // Handler, resolvers...
		fx.Invoke(
			Register,
			InitServer,
		), // Start server
	)
	app.Run()
}
