package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserGetAll(t *testing.T) {
	require := require.New(t)

	repo, err := NewUserRepoImpl()
	require.Nil(err)

	users, err := repo.GetAll(context.Background())
	require.Nil(err)
	require.NotEmpty(users)
}
