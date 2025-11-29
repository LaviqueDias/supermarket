package repositories

import (
	"database/sql"
	"github.com/LaviqueDias/supermarket/internal/domain/models"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) GetCartByClientID(clientID int) (int, error) {
	var cartID int
	err := r.db.QueryRow("SELECT id FROM cart WHERE client_id = ?", clientID).Scan(&cartID)
	return cartID, err
}

func (r *CartRepository) AddItem(cartID, productID int, quantity int, price float64) error {
	var existingID int
	err := r.db.QueryRow("SELECT id FROM cart_item WHERE cart_id = ? AND product_id = ?", cartID, productID).Scan(&existingID)

	if err == sql.ErrNoRows {
		query := `INSERT INTO cart_item (cart_id, product_id, quantity, unit_price) VALUES (?, ?, ?, ?)`
		_, err = r.db.Exec(query, cartID, productID, quantity, price)
	} else {
		query := `UPDATE cart_item SET quantity = quantity + ? WHERE id = ?`
		_, err = r.db.Exec(query, quantity, existingID)
	}
	return err
}

func (r *CartRepository) GetCartItems(clientID int) ([]map[string]interface{}, float64, error) {
	query := `
		SELECT ci.id, ci.cart_id, ci.product_id, ci.quantity, ci.unit_price, 
		       p.name, p.description
		FROM cart_item ci
		JOIN cart c ON ci.cart_id = c.id
		JOIN product p ON ci.product_id = p.id
		WHERE c.client_id = ?
	`
	rows, err := r.db.Query(query, clientID)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []map[string]interface{}
	var total float64

	for rows.Next() {
		var ci models.CartItem
		var productName, productDesc string
		err := rows.Scan(&ci.ID, &ci.CartID, &ci.ProductID, &ci.Quantity, &ci.UnitPrice, &productName, &productDesc)
		if err != nil {
			continue
		}

		subtotal := ci.UnitPrice * float64(ci.Quantity)
		total += subtotal

		items = append(items, map[string]interface{}{
			"id":          ci.ID,
			"product_id":  ci.ProductID,
			"name":        productName,
			"description": productDesc,
			"quantity":    ci.Quantity,
			"unit_price":  ci.UnitPrice,
			"subtotal":    subtotal,
		})
	}
	return items, total, nil
}

func (r *CartRepository) RemoveItem(itemID int) error {
	_, err := r.db.Exec("DELETE FROM cart_item WHERE id = ?", itemID)
	return err
}
