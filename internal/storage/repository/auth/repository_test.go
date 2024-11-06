package auth

import (
	"context"
	"errors"
	"io"
	"log"
	"testing"

	gofakeit "github.com/brianvoe/gofakeit/v7"
	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgconn"
	require "github.com/stretchr/testify/require"
	bcrypt "golang.org/x/crypto/bcrypt"

	service_model "github.com/MGomed/auth/internal/model"
	repo_errors "github.com/MGomed/auth/internal/storage/repository/errors"
	db_mock "github.com/MGomed/auth/pkg/client/db/mocks"
)

var (
	ctx    context.Context
	logger *log.Logger

	ctl     *gomock.Controller
	mockDB  *db_mock.MockDB
	mockDBC *db_mock.MockClient

	repo *repository

	errTest = errors.New("test")
)

func BeforeSuite(t *testing.T) {
	ctx = context.Background()
	logger = log.New(io.Discard, "", 0)

	ctl = gomock.NewController(t)
	mockDB = db_mock.NewMockDB(ctl)
	mockDBC = db_mock.NewMockClient(ctl)

	repo = &repository{dbc: mockDBC}

	t.Cleanup(ctl.Finish)
}

func TestCreate(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			user = &service_model.UserCreate{}

			expectedID int64 = 101

			r = &row{
				expectedID: expectedID,
			}
		)

		mockDBC.EXPECT().DB().Return(mockDB)
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), gomock.Any()).Return(r)

		id, err := repo.CreateUser(ctx, user)
		require.Equal(t, expectedID, id)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			user = &service_model.UserCreate{}

			r = &row{
				err: errTest,
			}
		)

		mockDBC.EXPECT().DB().Return(mockDB)
		mockDB.EXPECT().QueryRow(ctx, gomock.Any(), gomock.Any()).Return(r)

		_, err := repo.CreateUser(ctx, user)
		require.Equal(t, errors.Is(err, repo_errors.ErrQueryExecute), true)
	})

	t.Run("Rainy case (too long password)", func(t *testing.T) {
		var (
			user = &service_model.UserCreate{
				Password: []byte(gofakeit.Password(true, true, true, true, true, 73)),
			}
		)

		_, err := repo.CreateUser(ctx, user)
		require.Equal(t, errors.Is(err, bcrypt.ErrPasswordTooLong), true)
	})
}

func TestDelete(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			id int64 = 101
		)

		mockDBC.EXPECT().DB().Return(mockDB)
		mockDB.EXPECT().Exec(ctx, gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, nil)

		_, err := repo.DeleteUser(ctx, id)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			id int64 = 101
		)

		mockDBC.EXPECT().DB().Return(mockDB)
		mockDB.EXPECT().Exec(ctx, gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, errTest)

		_, err := repo.DeleteUser(ctx, id)
		require.Equal(t, errors.Is(err, repo_errors.ErrQueryExecute), true)
	})
}

func TestGet(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			id int64 = 101
		)

		mockDBC.EXPECT().DB().Return(mockDB)
		mockDB.EXPECT().ScanOne(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		_, err := repo.GetUser(ctx, id)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			id int64 = 101
		)

		mockDBC.EXPECT().DB().Return(mockDB)
		mockDB.EXPECT().ScanOne(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(errTest)

		_, err := repo.GetUser(ctx, id)
		require.Equal(t, errors.Is(err, repo_errors.ErrQueryExecute), true)
	})
}

func TestUpdate(t *testing.T) {
	BeforeSuite(t)

	t.Run("Sunny case", func(t *testing.T) {
		var (
			id   int64 = 101
			name       = "Alex"
			role       = service_model.RoleNames[1]

			user = &service_model.UserUpdate{
				ID:   id,
				Name: &name,
				Role: &role,
			}
		)

		mockDBC.EXPECT().DB().Return(mockDB)
		mockDB.EXPECT().Exec(ctx, gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, nil)

		_, err := repo.UpdateUser(ctx, user)
		require.Equal(t, nil, err)
	})

	t.Run("Rainy case", func(t *testing.T) {
		var (
			id   int64 = 101
			name       = "Alex"
			role       = service_model.RoleNames[1]

			user = &service_model.UserUpdate{
				ID:   id,
				Name: &name,
				Role: &role,
			}
		)

		mockDBC.EXPECT().DB().Return(mockDB)
		mockDB.EXPECT().Exec(ctx, gomock.Any(), gomock.Any()).Return(pgconn.CommandTag{}, errTest)

		_, err := repo.UpdateUser(ctx, user)
		require.Equal(t, errors.Is(err, repo_errors.ErrQueryExecute), true)
	})

	t.Run("Rainy case (empty update)", func(t *testing.T) {
		var (
			id int64 = 101

			user = &service_model.UserUpdate{
				ID: id,
			}
		)

		_, err := repo.UpdateUser(ctx, user)
		require.Equal(t, err != nil, true)
	})
}

type row struct {
	expectedID int64
	err        error
}

func (r row) Scan(dest ...interface{}) error {
	if r.err != nil {
		return r.err
	}

	val, _ := dest[0].(*int64)
	*val = r.expectedID

	return nil
}
