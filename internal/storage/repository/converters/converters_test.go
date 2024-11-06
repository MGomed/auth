package converters

import (
	"database/sql"
	"testing"
	"time"

	assert "github.com/stretchr/testify/assert"

	repo_model "github.com/MGomed/auth/internal/storage/repository/model"
)

func TestToUserFromRepo(t *testing.T) {
	t.Run("Sunny case", func(t *testing.T) {
		var info = &repo_model.UserInfo{
			ID:        1,
			Name:      "Alex",
			CreatedAt: time.Now(),
			UpdatedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		}

		user := ToUserFromRepo(info)
		assert.Equal(t, info.ID, user.ID)
		assert.Equal(t, info.Name, user.Name)
		assert.Equal(t, info.CreatedAt, user.CreatedAt)
		assert.Equal(t, info.UpdatedAt.Time, *user.UpdatedAt)
	})

	t.Run("Sunny case (not updated)", func(t *testing.T) {
		var info = &repo_model.UserInfo{
			ID:        1,
			Name:      "Alex",
			CreatedAt: time.Now(),
		}

		user := ToUserFromRepo(info)
		assert.Equal(t, info.ID, user.ID)
		assert.Equal(t, info.Name, user.Name)
		assert.Equal(t, info.CreatedAt, user.CreatedAt)
		assert.Equal(t, (*time.Time)(nil), user.UpdatedAt)
	})

	t.Run("Rainy case (nil)", func(t *testing.T) {
		user := ToUserFromRepo(nil)
		assert.Equal(t, user == nil, true)
	})
}
