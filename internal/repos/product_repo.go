package repos

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/roman-wb/crud-products/internal/models"
)

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (s *ProductRepo) All(ctx context.Context) (*[]models.Product, error) {
	var products []models.Product
	sql := `SELECT * FROM products ORDER BY id`
	err := pgxscan.Select(ctx, s.db, &products, sql)
	if err != nil {
		return nil, err
	}
	return &products, nil
}

func (s *ProductRepo) Find(ctx context.Context, id int) (*models.Product, error) {
	var product models.Product
	sql := `SELECT * FROM products WHERE id = $1 LIMIT 1`
	err := pgxscan.Get(ctx, s.db, &product, sql, id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *ProductRepo) Create(ctx context.Context, product *models.Product) error {
	sql := `INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id`
	return pgxscan.Get(ctx, s.db, &product.Id, sql, product.Name, product.Price)
}

func (s *ProductRepo) Update(ctx context.Context, product *models.Product) error {
	sql := `UPDATE products SET name = $1, price = $2 WHERE id = $3`
	_, err := s.db.Exec(ctx, sql, product.Name, product.Price, product.Id)
	return err
}

func (s *ProductRepo) Destroy(ctx context.Context, id int) error {
	sql := `DELETE FROM products WHERE id = $1`
	_, err := s.db.Exec(ctx, sql, id)
	return err
}
