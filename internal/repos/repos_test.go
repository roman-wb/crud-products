package repos

import (
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

func Test_NewRepos(t *testing.T) {
	db := &pgxpool.Pool{}
	repos := NewRepos(db)

	assert.NotEqual(t, nil, 1)

	assert.NotNil(t, repos.Product)
	assert.Equal(t, db, repos.Product.db)
}
