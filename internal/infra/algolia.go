//go:generate mockgen -package=infra -self_package=com.capturetweet/internal/infra -destination=algolia_mock.go . IndexInterface
package infra

import (
	"errors"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"os"
)

var (
	ErrInvalidSearchConfig = errors.New("invalid search client config, check your env_variables")
)

func NewIndex() (IndexInterface, error) {
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

type IndexInterface interface {
	Delete(opts ...interface{}) (res search.DeleteTaskRes, err error)

	// Indexing
	GetObject(objectID string, object interface{}, opts ...interface{}) error
	SaveObject(object interface{}, opts ...interface{}) (res search.SaveObjectRes, err error)

	// Searching
	Search(query string, opts ...interface{}) (res search.QueryRes, err error)
}
