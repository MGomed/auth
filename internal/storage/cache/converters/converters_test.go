package converters

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/assert"

	service_model "github.com/MGomed/auth/internal/model"
	cache_model "github.com/MGomed/auth/internal/storage/cache/model"
)

func TestToUserInfoFromCache(t *testing.T) {
	t.Run("Sunny case", func(t *testing.T) {
		var (
			now        = time.Now()
			id   int64 = 101
			info       = cache_model.UserInfo{
				Name:      "Alex",
				Email:     "Alex@main.com",
				CreatedAt: now.UnixNano(),
				UpdatedAt: now.UnixNano(),
			}
		)

		user := ToUserInfoFromCache(id, &info)
		assert.Equal(t, id, user.ID)
		assert.Equal(t, info.Name, user.Name)
		assert.Equal(t, info.Email, user.Email)
		assert.Equal(t, info.CreatedAt, user.CreatedAt.UnixNano())
		assert.Equal(t, info.UpdatedAt, user.UpdatedAt.UnixNano())
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			id   int64 = 101
			info *cache_model.UserInfo
		)

		user := ToUserInfoFromCache(id, info)
		assert.Equal(t, user == nil, true)
	})
}

func TestToUserInfoFromService(t *testing.T) {
	t.Run("Sunny case", func(t *testing.T) {
		var (
			now = time.Now().UTC()

			info = service_model.UserInfo{
				Name:      "Alex",
				Email:     "Alex@main.com",
				CreatedAt: now,
				UpdatedAt: &now,
			}
		)

		user := ToUserInfoFromService(&info)
		assert.Equal(t, info.Name, user.Name)
		assert.Equal(t, info.Email, user.Email)
		assert.Equal(t, info.CreatedAt, time.Unix(0, user.CreatedAt).UTC())
		assert.Equal(t, *info.UpdatedAt, time.Unix(0, user.UpdatedAt).UTC())
	})

	t.Run("Rainy case", func(t *testing.T) {
		var info *service_model.UserInfo

		user := ToUserInfoFromService(info)
		assert.Equal(t, user == nil, true)
	})
}
