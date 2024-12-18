package userapiimpl

import (
	"context"
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	require "github.com/stretchr/testify/require"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	api_errors "github.com/MGomed/auth/internal/api/errors"
	converters "github.com/MGomed/auth/internal/converters"
	service_mock "github.com/MGomed/auth/internal/service/mocks"
	user_api "github.com/MGomed/auth/pkg/user_api"
)

var (
	ctx context.Context

	ctl         *gomock.Controller
	mockService *service_mock.MockUserService

	api *UserAPI

	errTest = errors.New("test")
)

func BeforeSuite(t *testing.T) {
	ctx = context.Background()

	ctl = gomock.NewController(t)
	mockService = service_mock.NewMockUserService(ctl)

	api = &UserAPI{service: mockService}

	t.Cleanup(ctl.Finish)
}

func TestCreate(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			user = &user_api.UserCreate{
				Name:            "Alex",
				Email:           "Alex@mail.com",
				Password:        "Password",
				PasswordConfirm: "Password",
			}

			expectedID = int64(101)
		)

		mockService.EXPECT().Create(gomock.Any(), converters.ToUserCreateFromAPI(user)).Return(expectedID, nil)

		resp, err := api.Create(ctx, &user_api.CreateRequest{User: user})

		require.Equal(t, resp.Id, expectedID)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			user = &user_api.UserCreate{
				Name:            "Alex",
				Email:           "Alex@mail.com",
				Password:        "Password",
				PasswordConfirm: "Password",
			}
		)

		mockService.EXPECT().Create(gomock.Any(), converters.ToUserCreateFromAPI(user)).Return(int64(0), errTest)

		_, err := api.Create(ctx, &user_api.CreateRequest{User: user})

		require.Equal(t, errTest, err)
	})

	t.Run("Rainy case (password mismatch)", func(t *testing.T) {
		var (
			user = &user_api.UserCreate{
				Name:            "Alex",
				Email:           "Alex@mail.com",
				Password:        "pass1",
				PasswordConfirm: "pass2",
			}
		)

		_, err := api.Create(ctx, &user_api.CreateRequest{User: user})
		require.Equal(t, api_errors.ErrPasswordMismatch, err)
	})
}

func TestDelete(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			id = int64(101)
		)

		mockService.EXPECT().Delete(gomock.Any(), id).Return(nil)

		_, err := api.Delete(ctx, &user_api.DeleteRequest{Id: id})
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			id = int64(101)
		)

		mockService.EXPECT().Delete(gomock.Any(), id).Return(errTest)

		_, err := api.Delete(ctx, &user_api.DeleteRequest{Id: id})
		require.Equal(t, errTest, err)
	})
}

func TestGet(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			id = int64(101)
		)

		mockService.EXPECT().Get(gomock.Any(), id).Return(nil, nil)

		_, err := api.Get(ctx, &user_api.GetRequest{Id: id})
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			id = int64(101)
		)

		mockService.EXPECT().Get(gomock.Any(), id).Return(nil, errTest)

		_, err := api.Get(ctx, &user_api.GetRequest{Id: id})
		require.Equal(t, errTest, err)
	})
}

func TestUpdate(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			id   int64 = 101
			user       = &user_api.UserUpdate{
				Name: &wrapperspb.StringValue{
					Value: "Alex",
				},
			}
		)

		mockService.EXPECT().Update(gomock.Any(), id, converters.ToUserUpdateFromAPI(user)).Return(nil)

		_, err := api.Update(ctx, &user_api.UpdateRequest{Id: id, User: user})
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			id   int64 = 101
			user       = &user_api.UserUpdate{
				Name: &wrapperspb.StringValue{
					Value: "Alex",
				},
				Role: user_api.Role_UNKNOWN,
			}
		)

		mockService.EXPECT().Update(gomock.Any(), id, converters.ToUserUpdateFromAPI(user)).Return(errTest)

		_, err := api.Update(ctx, &user_api.UpdateRequest{Id: id, User: user})
		require.Equal(t, errTest, err)
	})
}
