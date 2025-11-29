package repositories

import (
	"database/sql"
	"github.com/LaviqueDias/supermarket/internal/domain/models"
)

type PromotionRepository struct {
	db *sql.DB
}

func NewPromotionRepository(db *sql.DB) *PromotionRepository {
	return &PromotionRepository{db: db}
}

func (r *PromotionRepository) Create(p *models.Promotion) error {
	query := `INSERT INTO promotion (name, discount_percent, start_date, end_date, active) VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, p.Name, p.DiscountPercent, p.StartDate, p.EndDate, p.Active)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	p.ID = int(id)
	return nil
}

func (r *PromotionRepository) FindAll() ([]models.Promotion, error) {
	query := "SELECT id, name, discount_percent, start_date, end_date, active, created_at, updated_at FROM promotion"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var promotions []models.Promotion
	for rows.Next() {
		var p models.Promotion
		err := rows.Scan(&p.ID, &p.Name, &p.DiscountPercent, &p.StartDate, &p.EndDate, &p.Active, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			continue
		}
		promotions = append(promotions, p)
	}
	return promotions, nil
}

func (r *PromotionRepository) AddProduct(promotionID, productID int) error {
	query := `INSERT INTO promotion_products (promotion_id, product_id) VALUES (?, ?)`
	_, err := r.db.Exec(query, promotionID, productID)
	return err
}
