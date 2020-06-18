package search

import (
	"com.capturetweet/internal/infra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func _TestService_PutSearch(t *testing.T) {

	index, err := infra.NewIndex()
	require.NoError(t, err)
	require.NotNil(t, index)

	service := NewService(index)

	// err = service.Put(uuid.UUIDv4(), "test text", "rayyildiz")
	// require.NoError(t, err)

	searchModels, err := service.Search("test", 20)
	require.NoError(t, err)
	require.NotNil(t, searchModels)
	assert.Equal(t, 4, len(searchModels))
}
