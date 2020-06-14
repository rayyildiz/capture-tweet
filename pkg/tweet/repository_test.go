package tweet

import (
	"com.capturetweet/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepository_Store(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	db.AutoMigrate(&model.User{}, &model.Tweet{})

	repo := NewRepository(db)
	err = repo.Store("id1", "full text", "en", "userId1", 1, 1, nil)
	require.NoError(t, err)
}
