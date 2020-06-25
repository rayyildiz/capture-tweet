package main

import (
	"com.capturetweet/internal/infra"
	"com.capturetweet/pkg/graph"
	"com.capturetweet/pkg/search"
	"com.capturetweet/pkg/tweet"
	"com.capturetweet/pkg/user"
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	godotenv.Load()
}

func main() {
	start := time.Now()
	logger := infra.NewLogger()
	ensureNotNil(logger, "zap:logger")

	tweetColl, err := infra.NewTweetCollection()
	ensureNoError(err, "twitter:docstore collection")
	defer tweetColl.Close()

	userColl, err := infra.NewUserCollection()
	ensureNoError(err, "user:docstore collection")
	defer userColl.Close()

	topic, err := infra.NewTopic(os.Getenv("TOPIC_CAPTURE"))
	ensureNoError(err, "pubsub topic capture")
	defer topic.Shutdown(context.Background())

	searchIndexer, err := infra.NewIndex()
	ensureNoError(err, "search index, algolia")

	twitterApi := infra.NewTwitterClient()
	ensureNotNil(twitterApi, "anaconda:twitter client")

	searchService := search.NewService(searchIndexer)
	ensureNotNil(searchService, "search:NewService")

	userService := user.NewService(user.NewRepository(userColl), logger)
	ensureNotNil(userService, "user:NewService")

	tweetService := tweet.NewService(tweet.NewRepository(tweetColl), searchService, userService, twitterApi, logger, topic)
	ensureNotNil(tweetService, "tweet:NewService")

	rootResolver := graph.NewResolver()
	ensureNotNil(rootResolver, "graph:NewResolver")
	graph.InitService(logger, tweetService, userService)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: rootResolver}))
	srv.Use(infra.ZapLogger{Log: logger})

	if os.Getenv("GRAPHQL_ENABLE_PLAYGROUND") == "true" {
		http.Handle("/", playground.Handler("GraphQL playground", "/api/query"))
	}
	http.Handle("/api/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	diff := time.Now().Sub(start)
	logger.Info("initialized objects", zap.Duration("elapsed", diff))

	err = http.ListenAndServe(":"+port, nil)
	ensureNoError(err, "http:ListenAndServe, port :"+port)
}

func ensureNoError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s, %v", msg, err)
	}
}

func ensureNotNil(obj interface{}, msg string) {
	if obj == nil {
		log.Fatalf("object is nil, %s", msg)
	}
}
