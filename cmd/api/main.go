package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"com.capturetweet/internal/infra"
	"com.capturetweet/pkg/graph"
	"com.capturetweet/pkg/search"
	"com.capturetweet/pkg/tweet"
	"com.capturetweet/pkg/user"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go.uber.org/zap"
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

	mux := http.NewServeMux()

	if os.Getenv("GRAPHQL_ENABLE_PLAYGROUND") == "true" {
		mux.Handle("/", playground.Handler("GraphQL playground", "/api/query"))
	}
	mux.Handle("/api/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://beta.capturetweet.com", "http://localhost:3000"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	}).Handler(mux)

	// handler := cors.AllowAll().Handler(mux)

	diff := time.Now().Sub(start)
	logger.Info("initialized objects", zap.Duration("elapsed", diff))

	err = http.ListenAndServe(":"+port, handler)
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
