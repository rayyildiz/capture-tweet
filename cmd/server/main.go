package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"capturetweet.com/internal/infra"
	"capturetweet.com/pkg/content"
	"capturetweet.com/pkg/resolver"
	"capturetweet.com/pkg/search"
	"capturetweet.com/pkg/tweet"
	"capturetweet.com/pkg/user"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("can't load .env file, %v", err)
	}
}

func main() {
	if err := Run(); err != nil {
		log.Fatalf("%v", err)
	}
}

func Run() error {
	defer sentry.Flush(time.Second * 2)

	start := time.Now()

	tweetColl := infra.NewTweetCollection()
	defer tweetColl.Close()

	userColl := infra.NewUserCollection()
	defer userColl.Close()

	contactUsColl := infra.NewContactUsCollection()
	defer contactUsColl.Close()

	topic := infra.NewTopic(os.Getenv("TOPIC_CAPTURE"))
	defer topic.Shutdown(context.Background())

	searchIndexer := infra.NewIndex()

	twitterApi := infra.NewTwitterClient()
	if twitterApi == nil {
		return fmt.Errorf("anaconda:twitter client is nil")
	}

	searchService := search.NewService(searchIndexer)
	if searchService == nil {
		return fmt.Errorf("search:NewService is nil")
	}

	userService := user.NewService(user.NewRepository(userColl))
	if userService == nil {
		return fmt.Errorf("user:NewService is nil")
	}

	tweetService := tweet.NewService(tweet.NewRepository(tweetColl), searchService, userService, twitterApi, topic)
	if tweetService == nil {
		return fmt.Errorf("tweet:NewService is nil")
	}

	contentService := content.NewService(content.NewRepository(contactUsColl))
	if contentService == nil {
		return fmt.Errorf("content service is nil")
	}

	rootResolver := resolver.NewResolver()
	if rootResolver == nil {
		return fmt.Errorf("graph:NewResolver is nil")
	}
	resolver.InitService(tweetService, userService, contentService)

	srv := handler.NewDefaultServer(resolver.NewExecutableSchema(resolver.Config{Resolvers: rootResolver}))

	mux := http.DefaultServeMux

	if os.Getenv("GRAPHQL_ENABLE_PLAYGROUND") == "true" {
		mux.Handle("GET /api/docs", playground.Handler("GraphQL playground", "/api/query"))
	}
	mux.Handle("ANY /api/query", infra.VersionHandler(srv))

	h := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://capturetweet.com", "https://www.capturetweet.com", "http://localhost:3000", "http://localhost:4000"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}).Handler(mux)

	slog.Info("initialized objects", slog.Duration("elapsed", time.Since(start).Round(time.Millisecond)))

	port := infra.Port()
	slog.Info("api server is starting at port", slog.String("port", port))
	if err := http.ListenAndServe(":"+port, h); err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}

	return nil
}
