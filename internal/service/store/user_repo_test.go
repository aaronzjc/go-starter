package store

import (
	"context"
	"go-starter/internal/domain/model"
	"go-starter/test"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserGetAll(t *testing.T) {
	require := require.New(t)

	test.SetupTestDb(t, model.DB_DEMO)

	repo, err := NewUserRepoImpl()
	require.Nil(err)

	users, err := repo.GetAll(context.Background())
	require.Nil(err)
	require.NotEmpty(users)
}
