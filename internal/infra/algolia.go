//go:generate mockgen -package=infra -self_package=com.capturetweet/internal/infra -destination=algolia_mock.go . IndexInterface
package infra

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"log"
	"os"
)

func NewIndex() IndexInterface {
	secret := os.Getenv("ALGOLIA_SECRET")
	clientId := os.Getenv("ALGOLIA_CLIENT_ID")
	indexName := os.Getenv("ALGOLIA_INDEX")
	if len(secret) == 0 || len(clientId) == 0 || len(indexName) == 0 {
		log.Fatalf("check your algolia system env variables: ALGOLIA_SECRET,ALGOLIA_CLIENT_ID and ALGOLIA_INDEX")
	}

	client := search.NewClient(clientId, secret)
	index := client.InitIndex(indexName)
	return index
}

type IndexInterface interface {
	Delete(opts ...interface{}) (res search.DeleteTaskRes, err error)

	// GetObject return single object
	GetObject(objectID string, object interface{}, opts ...interface{}) error
	SaveObject(object interface{}, opts ...interface{}) (res search.SaveObjectRes, err error)

	// Search return for searching api.
	Search(query string, opts ...interface{}) (res search.QueryRes, err error)
}
