package repositories

import (
	"database/sql"
	"github.com/LaviqueDias/supermarket/internal/domain/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(p *models.Product) error {
	query := `INSERT INTO product (name, description, price, stock_quantity) VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(query, p.Name, p.Description, p.Price, p.StockQuantity)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	p.ID = int(id)
	return nil
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	query := "SELECT id, name, description, price, stock_quantity, created_at, updated_at FROM product"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.StockQuantity, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			continue
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) FindByID(id int) (*models.Product, error) {
	query := "SELECT id, name, description, price, stock_quantity, created_at, updated_at FROM product WHERE id = ?"
	var p models.Product
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.StockQuantity, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Update(p *models.Product) error {
	query := `UPDATE product SET name=?, description=?, price=?, stock_quantity=? WHERE id=?`
	_, err := r.db.Exec(query, p.Name, p.Description, p.Price, p.StockQuantity, p.ID)
	return err
}

func (r *ProductRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM product WHERE id=?", id)
	return err
}
