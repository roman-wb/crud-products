package repos

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/roman-wb/crud-products/internal/models"
	"github.com/roman-wb/crud-products/pkg/test"
	"github.com/stretchr/testify/require"
)

func Test_NewProductRepo(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	db := &pgxpool.Pool{}
	repo := NewProductRepo(db)

	require.Equal(t, db, repo.db)
}

func Test_ProductRepo_All(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	db := test.Setup()
	repo := NewProductRepo(db)

	t.Run("Exist products", func(t *testing.T) {
		defer test.Truncate()

		wantProducts := []models.Product{
			{
				Id:    1,
				Name:  "Test 1",
				Price: 100.99,
			},
			{
				Id:    2,
				Name:  "Test 2",
				Price: 0,
			},
		}

		sql := `INSERT INTO products (id, name, price) VALUES (1, 'Test 1', 100.99), (2, 'Test 2', 0)`
		_, err := db.Exec(context.Background(), sql)
		require.Nil(t, err)

		gotProducts, err := repo.All(context.Background())
		require.Nil(t, err)

		require.Equal(t, len(wantProducts), len(*gotProducts))
		for i, tc := range wantProducts {
			require.Equal(t, tc, (*gotProducts)[i])
		}
	})

	t.Run("Not exist products", func(t *testing.T) {
		defer test.Truncate()

		goytProducts, err := repo.All(context.Background())
		require.Nil(t, err)

		require.Equal(t, 0, len(*goytProducts))
	})
}

func Test_ProductRepo_Find(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	db := test.Setup()
	repo := NewProductRepo(db)

	t.Run("Exist product", func(t *testing.T) {
		defer test.Truncate()

		wantProduct := models.Product{Id: 1, Name: "Test 1", Price: 100.99}

		sql := `INSERT INTO products (id, name, price) VALUES (1, 'Test 1', 100.99), (2, 'Test 2', 0)`
		_, err := db.Exec(context.Background(), sql)
		require.Nil(t, err)

		gotProduct, err := repo.Find(context.Background(), 1)
		require.Nil(t, err)

		require.Equal(t, wantProduct, *gotProduct)
	})

	t.Run("Not exist product", func(t *testing.T) {
		defer test.Truncate()

		sql := `INSERT INTO products (id, name, price) VALUES (1, 'Test 1', 100.99), (2, 'Test 2', 0)`
		_, err := db.Exec(context.Background(), sql)
		require.Nil(t, err)

		gotProduct, err := repo.Find(context.Background(), 10)

		require.Nil(t, gotProduct)
		require.Error(t, pgx.ErrNoRows, err)
	})
}

func Test_Create(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	db := test.Setup()
	defer test.Truncate()

	repo := NewProductRepo(db)
	product := &models.Product{Id: 1, Name: "Test 1", Price: 100.99}

	gotErr := repo.Create(context.Background(), product)

	require.Nil(t, gotErr)
}

func Test_Update(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	db := test.Setup()
	defer test.Truncate()

	repo := NewProductRepo(db)

	wantProduct1 := &models.Product{Id: 1, Name: "Test 1 - updated", Price: 1999.99}
	wantProduct2 := &models.Product{Id: 2, Name: "Test 2", Price: 0}

	sql := `INSERT INTO products (id, name, price) VALUES (1, 'Test 1', 100.99), (2, 'Test 2', 0)`
	_, err := db.Exec(context.Background(), sql)
	require.Nil(t, err)

	err = repo.Update(context.Background(), wantProduct1)
	require.Nil(t, err)

	gotProducts, gotErr := repo.All(context.Background())
	require.Nil(t, gotErr)

	require.Equal(t, (*gotProducts)[0], *wantProduct1)
	require.Equal(t, (*gotProducts)[1], *wantProduct2)
}

func Test_Destroy(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	db := test.Setup()
	defer test.Truncate()

	repo := NewProductRepo(db)

	wantProduct := models.Product{Id: 1, Name: "Test 1", Price: 100.99}

	sql := `INSERT INTO products (id, name, price) VALUES (1, 'Test 1', 100.99), (2, 'Test 2', 0)`
	_, err := db.Exec(context.Background(), sql)
	require.Nil(t, err)

	err = repo.Destroy(context.Background(), 2)
	require.Nil(t, err)

	gotProducts, err := repo.All(context.Background())
	require.Nil(t, err)

	require.Equal(t, 1, len(*gotProducts))
	require.Equal(t, wantProduct, (*gotProducts)[0])
}
