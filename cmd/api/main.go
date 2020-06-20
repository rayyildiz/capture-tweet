package main

import (
	"com.capturetweet/internal/infra"
	"com.capturetweet/pkg/graph"
	"com.capturetweet/pkg/search"
	"com.capturetweet/pkg/tweet"
	"com.capturetweet/pkg/user"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	godotenv.Load()
}

func main() {
	logger := infra.NewLogger()
	ensureNotNil(logger, "zap:logger")

	tweetColl, err := infra.NewTweetCollection()
	ensureNoError(err, "twitter:docstore collection")

	userColl, err := infra.NewUserCollection()
	ensureNoError(err, "user:docstore collection")

	searchIndexer, err := infra.NewIndex()
	ensureNoError(err, "search index, algolia")

	searchService := search.NewService(searchIndexer)
	ensureNotNil(searchService, "search:NewService")

	userService := user.NewService(user.NewRepository(userColl))
	ensureNotNil(userService, "user:NewService")

	tweetService := tweet.NewService(tweet.NewRepository(tweetColl), searchService)
	ensureNotNil(tweetService, "tweet:NewService")

	rootResolver := graph.NewResolver()
	ensureNotNil(rootResolver, "graph:NewResolver")
	graph.InitService(tweetService, searchService, userService)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: rootResolver}))

	if os.Getenv("GRAPHQL_ENABLE_PLAYGROUND") == "true" {
		http.Handle("/", playground.Handler("GraphQL playground", "/api/query"))
	}
	http.Handle("/api/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

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
