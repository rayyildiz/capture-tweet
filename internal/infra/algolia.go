package infra

import (
	"errors"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"os"
)

var (
	ErrInvalidSearchConfig = errors.New("invalid search client config, check your env_variables")
)

func NewIndex() (*search.Index, error) {
	secret := os.Getenv("ALGOLIA_SECRET")
	clientId := os.Getenv("ALGOLIA_CLIENT_ID")
	indexName := os.Getenv("ALGOLIA_INDEX")
	if len(secret) == 0 || len(clientId) == 0 || len(indexName) == 0 {
		return nil, ErrInvalidSearchConfig
	}

	client := search.NewClient(clientId, secret)
	index := client.InitIndex(indexName)
	return index, nil
}
