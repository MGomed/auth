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
	repo_mock "github.com/MGomed/auth/internal/storage/repository/mocks"
)

var (
	ctx    context.Context
	logger *log.Logger

	ctl *gomock.Controller

	mockRepo *repo_mock.MockRepository

	serv *service

	errTest = errors.New("test")
)

func BeforeSuite(t *testing.T) {
	ctx = context.Background()
	logger = log.New(io.Discard, "", 0)

	ctl = gomock.NewController(t)
	mockRepo = repo_mock.NewMockRepository(ctl)

	serv = &service{logger: logger, repo: mockRepo}

	t.Cleanup(ctl.Finish)
}

func TestCreate(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			user service_model.UserCreate

			expectedID int64 = 101
		)

		mockRepo.EXPECT().CreateUser(ctx, &user).Return(expectedID, nil)

		id, err := serv.Create(ctx, &user)

		require.Equal(t, expectedID, id)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			user service_model.UserCreate
		)

		mockRepo.EXPECT().CreateUser(ctx, &user).Return(int64(0), errTest)

		_, err := serv.Create(ctx, &user)

		require.Equal(t, errTest, err)
	})
}

func TestGet(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			id int64 = 101
		)

		mockRepo.EXPECT().GetUser(ctx, id).Return(&service_model.UserInfo{}, nil)

		_, err := serv.Get(ctx, id)

		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			id int64 = 101
		)

		mockRepo.EXPECT().GetUser(ctx, id).Return(nil, errTest)

		_, err := serv.Get(ctx, id)

		require.Equal(t, errTest, err)
	})
}

func TestUpdate(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			user = &service_model.UserUpdate{
				ID: 1,
			}
		)

		mockRepo.EXPECT().UpdateUser(ctx, user).Return(int64(0), nil)

		err := serv.Update(ctx, user)

		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			user = &service_model.UserUpdate{
				ID: 1,
			}
		)

		mockRepo.EXPECT().UpdateUser(ctx, user).Return(int64(0), errTest)

		err := serv.Update(ctx, user)

		require.Equal(t, errTest, err)
	})
}

func TestDelete(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			id int64 = 101
		)

		mockRepo.EXPECT().DeleteUser(ctx, id).Return(int64(0), nil)

		err := serv.Delete(ctx, id)

		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			id int64 = 101
		)

		mockRepo.EXPECT().DeleteUser(ctx, id).Return(int64(0), errTest)

		err := serv.Delete(ctx, id)

		require.Equal(t, errTest, err)
	})
}
