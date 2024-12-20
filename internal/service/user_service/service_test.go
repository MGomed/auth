package userservice

import (
	"context"
	"errors"
	"io"
	"log"
	"testing"

	gomock "github.com/golang/mock/gomock"
	require "github.com/stretchr/testify/require"

	service_model "github.com/MGomed/auth/internal/model"
	cache_errors "github.com/MGomed/auth/internal/storage/cache/errors"
	storage_mock "github.com/MGomed/auth/internal/storage/mocks"
	db_mock "github.com/MGomed/common/client/db/mocks"
)

var (
	ctx    context.Context
	logger *log.Logger

	ctl *gomock.Controller

	mockRepo     *storage_mock.MockRepository
	mockCache    *storage_mock.MockCache
	mockTxMnager *db_mock.MockTxManager
	mockMsgBus   *storage_mock.MockMessageBus

	serv *service

	errTest = errors.New("test")
)

func BeforeSuite(t *testing.T) {
	ctx = context.Background()
	logger = log.New(io.Discard, "", 0)

	ctl = gomock.NewController(t)
	mockRepo = storage_mock.NewMockRepository(ctl)
	mockCache = storage_mock.NewMockCache(ctl)
	mockTxMnager = db_mock.NewMockTxManager(ctl)
	mockMsgBus = storage_mock.NewMockMessageBus(ctl)

	serv = &service{
		logger:    logger,
		repo:      mockRepo,
		cache:     mockCache,
		txManager: mockTxMnager,
		msgBus:    mockMsgBus,
	}

	t.Cleanup(ctl.Finish)
}

func TestCreate(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		mockTxMnager.EXPECT().ReadCommitted(ctx, gomock.Any()).Return(nil)
		mockCache.EXPECT().CreateUser(ctx, gomock.Any(), gomock.Any()).Return(nil)
		mockMsgBus.EXPECT().SendMessage(ctx, gomock.Any()).Return(nil)

		_, err := serv.Create(ctx, &service_model.UserCreate{})
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case (failed to add user into database)", func(t *testing.T) {
		mockTxMnager.EXPECT().ReadCommitted(ctx, gomock.Any()).Return(errTest)

		_, err := serv.Create(ctx, &service_model.UserCreate{})
		require.ErrorIs(t, err, errTest)
	})

	t.Run("Rainy case (failed to add user into cache)", func(t *testing.T) {
		mockTxMnager.EXPECT().ReadCommitted(ctx, gomock.Any()).Return(nil)
		mockCache.EXPECT().CreateUser(ctx, gomock.Any(), gomock.Any()).Return(errTest)

		_, err := serv.Create(ctx, &service_model.UserCreate{})
		require.ErrorIs(t, err, errTest)
	})

	t.Run("Rainy case (failed send message to message bus)", func(t *testing.T) {
		mockTxMnager.EXPECT().ReadCommitted(ctx, gomock.Any()).Return(nil)
		mockCache.EXPECT().CreateUser(ctx, gomock.Any(), gomock.Any()).Return(nil)
		mockMsgBus.EXPECT().SendMessage(ctx, gomock.Any()).Return(errTest)

		_, err := serv.Create(ctx, &service_model.UserCreate{})
		require.ErrorIs(t, err, errTest)
	})
}

func TestGet(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case (user present in cache)", func(t *testing.T) {
		var id int64 = 101

		mockCache.EXPECT().GetUser(ctx, id).Return(&service_model.UserInfo{}, nil)

		_, err := serv.Get(ctx, id)
		require.Equal(t, nil, err)
	})

	t.Run("Sunny case (user don't present in cache)", func(t *testing.T) {
		var id int64 = 101

		mockCache.EXPECT().GetUser(ctx, id).Return(nil, cache_errors.ErrUserNotPresent)
		mockRepo.EXPECT().GetUser(ctx, id).Return(&service_model.UserInfo{}, nil)

		_, err := serv.Get(ctx, id)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var id int64 = 101

		mockCache.EXPECT().GetUser(ctx, id).Return(nil, cache_errors.ErrUserNotPresent)
		mockRepo.EXPECT().GetUser(ctx, id).Return(nil, errTest)

		_, err := serv.Get(ctx, id)
		require.ErrorIs(t, err, errTest)
	})
}

func TestUpdate(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		mockTxMnager.EXPECT().ReadCommitted(ctx, gomock.Any()).Return(nil)
		mockCache.EXPECT().CreateUser(ctx, gomock.Any(), gomock.Any()).Return(nil)

		err := serv.Update(ctx, 0, &service_model.UserUpdate{})

		require.Equal(t, nil, err)
	})

	t.Run("Rainy case (failed to update user into database)", func(t *testing.T) {
		mockTxMnager.EXPECT().ReadCommitted(ctx, gomock.Any()).Return(errTest)

		err := serv.Update(ctx, 0, &service_model.UserUpdate{})
		require.ErrorIs(t, err, errTest)
	})

	t.Run("Rainy case (failed to update user into cache)", func(t *testing.T) {
		mockTxMnager.EXPECT().ReadCommitted(ctx, gomock.Any()).Return(nil)
		mockCache.EXPECT().CreateUser(ctx, gomock.Any(), gomock.Any()).Return(errTest)

		err := serv.Update(ctx, 0, &service_model.UserUpdate{})
		require.ErrorIs(t, err, errTest)
	})
}

func TestDelete(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case (user present in cache)", func(t *testing.T) {
		var id int64 = 101

		mockRepo.EXPECT().DeleteUser(ctx, id).Return(int64(0), nil)
		mockCache.EXPECT().DeleteUser(ctx, id).Return(nil)
		mockMsgBus.EXPECT().SendMessage(ctx, gomock.Any()).Return(nil)

		err := serv.Delete(ctx, id)
		require.Equal(t, nil, err)
	})

	t.Run("Sunny case (user don't present in cache)", func(t *testing.T) {
		var id int64 = 101

		mockRepo.EXPECT().DeleteUser(ctx, id).Return(int64(0), nil)
		mockCache.EXPECT().DeleteUser(ctx, id).Return(cache_errors.ErrUserNotPresent)
		mockMsgBus.EXPECT().SendMessage(ctx, gomock.Any()).Return(nil)

		err := serv.Delete(ctx, id)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case (failed to delete user from database)", func(t *testing.T) {
		var id int64 = 101

		mockRepo.EXPECT().DeleteUser(ctx, id).Return(int64(0), errTest)

		err := serv.Delete(ctx, id)
		require.ErrorIs(t, err, errTest)
	})

	t.Run("Rainy case (failed to delete user from cache)", func(t *testing.T) {
		var id int64 = 101

		mockRepo.EXPECT().DeleteUser(ctx, id).Return(int64(0), nil)
		mockCache.EXPECT().DeleteUser(ctx, id).Return(errTest)

		err := serv.Delete(ctx, id)
		require.ErrorIs(t, err, errTest)
	})

	t.Run("Rainy case (failed send message to message bus)", func(t *testing.T) {
		var id int64 = 101

		mockRepo.EXPECT().DeleteUser(ctx, id).Return(int64(0), nil)
		mockCache.EXPECT().DeleteUser(ctx, id).Return(nil)
		mockMsgBus.EXPECT().SendMessage(ctx, gomock.Any()).Return(errTest)

		err := serv.Delete(ctx, id)
		require.ErrorIs(t, err, errTest)
	})
}
