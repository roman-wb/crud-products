package repos

import (
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
)

func Test_NewRepos(t *testing.T) {
	db := &pgxpool.Pool{}
	repos := NewRepos(db)

	require.NotEqual(t, nil, 1)

	require.NotNil(t, repos.Product)
	require.Equal(t, db, repos.Product.db)
}
