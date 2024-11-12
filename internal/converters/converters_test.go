package converters

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"

	service_model "github.com/MGomed/auth/internal/model"
	"github.com/MGomed/auth/pkg/user_api"
)

func TestToUserInfoFromService(t *testing.T) {
	t.Run("Sunny case", func(t *testing.T) {
		var (
			updatedTime = time.Now().UTC()

			info = &service_model.UserInfo{
				ID:        1,
				Name:      "Alex",
				CreatedAt: time.Now().UTC(),
				UpdatedAt: &updatedTime,
			}
		)

		user := ToUserInfoFromService(info)
		assert.Equal(t, info.ID, user.Id)
		assert.Equal(t, info.Name, user.Name)
		assert.Equal(t, info.CreatedAt, user.CreatedAt.AsTime())
		assert.Equal(t, *info.UpdatedAt, user.UpdatedAt.AsTime())
	})

	t.Run("Rainy case (nil)", func(t *testing.T) {
		user := ToUserInfoFromService(nil)
		assert.Equal(t, user == nil, true)
	})
}

func TestToUserCreateFromAPI(t *testing.T) {
	t.Run("Sunny case", func(t *testing.T) {
		var (
			create = &user_api.UserCreate{
				Name:  "Alex",
				Email: "Alex@mail.com",
			}
		)

		user := ToUserCreateFromAPI(create)
		assert.Equal(t, create.Name, user.Name)
		assert.Equal(t, create.Email, user.Email)
	})

	t.Run("Rainy case (nil)", func(t *testing.T) {
		user := ToUserCreateFromAPI(nil)
		assert.Equal(t, user == nil, true)
	})
}

func TestToUserUpdateFromAPI(t *testing.T) {
	t.Run("Sunny case", func(t *testing.T) {
		var (
			update = &user_api.UserUpdate{
				Id:   1,
				Name: wrapperspb.String("Alex"),
				Role: user_api.Role_USER,
			}
		)

		user := ToUserUpdateFromAPI(update)
		assert.Equal(t, update.Id, user.ID)
		assert.Equal(t, update.Name.Value, *user.Name)
		assert.Equal(t, user_api.Role_name[int32(update.Role)], *user.Role)
	})

	t.Run("Rainy case (nil)", func(t *testing.T) {
		user := ToUserUpdateFromAPI(nil)
		assert.Equal(t, user == nil, true)
	})
}
