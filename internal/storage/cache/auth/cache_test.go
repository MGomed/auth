package auth

import (
	"context"
	"errors"
	"io"
	"log"
	"testing"

	gomock "github.com/golang/mock/gomock"
	require "github.com/stretchr/testify/require"

	service_model "github.com/MGomed/auth/internal/model"
	cache_converters "github.com/MGomed/auth/internal/storage/cache/converters"
	cache_mock "github.com/MGomed/common/pkg/client/cache/mocks"
)

var (
	ctx    context.Context
	logger *log.Logger

	ctl        *gomock.Controller
	mockClient *cache_mock.MockRedisClient

	c *cacher

	errTest = errors.New("test")
)

func BeforeSuite(t *testing.T) {
	ctx = context.Background()
	logger = log.New(io.Discard, "", 0)

	ctl = gomock.NewController(t)
	mockClient = cache_mock.NewMockRedisClient(ctl)

	c = &cacher{client: mockClient}

	t.Cleanup(ctl.Finish)
}

func TestCreate(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			id int64 = 101

			user service_model.UserInfo
		)

		mockClient.EXPECT().
			HashSet(ctx, constructKey(id), cache_converters.ToUserInfoFromService(&user)).Return(nil)

		err := c.CreateUser(ctx, id, &user)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			id int64 = 101

			user service_model.UserInfo
		)

		mockClient.EXPECT().
			HashSet(ctx, constructKey(id), cache_converters.ToUserInfoFromService(&user)).Return(errTest)

		err := c.CreateUser(ctx, id, &user)
		require.Equal(t, errTest, err)
	})
}

func TestGet(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var id int64 = 101

		mockClient.EXPECT().
			HGetAll(ctx, constructKey(id), gomock.Any()).Return(nil)

		_, err := c.GetUser(ctx, id)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var id int64 = 101

		mockClient.EXPECT().
			HGetAll(ctx, constructKey(id), gomock.Any()).Return(errTest)

		_, err := c.GetUser(ctx, id)
		require.Equal(t, errTest, err)
	})
}

func TestDelete(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var id int64 = 101

		mockClient.EXPECT().
			Del(ctx, constructKey(id)).Return(nil)

		err := c.DeleteUser(ctx, id)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var id int64 = 101

		mockClient.EXPECT().
			Del(ctx, constructKey(id)).Return(errTest)

		err := c.DeleteUser(ctx, id)
		require.Equal(t, errTest, err)
	})
}
