package repos

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repos struct {
	Product *ProductRepo
}

func NewRepos(db *pgxpool.Pool) *Repos {
	return &Repos{
		Product: NewProductRepo(db),
	}
}
