package repositories

import (
	"database/sql"
	"github.com/LaviqueDias/supermarket/internal/domain/models"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) Create(c *models.Client) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO client (name, email, password_hash) VALUES (?, ?, ?)`
	result, err := tx.Exec(query, c.Name, c.Email, c.PasswordHash)
	if err != nil {
		tx.Rollback()
		return err
	}

	clientID, _ := result.LastInsertId()
	c.ID = int(clientID)

	_, err = tx.Exec("INSERT INTO cart (client_id) VALUES (?)", clientID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *ClientRepository) FindAll() ([]models.Client, error) {
	query := "SELECT id, name, email, created_at, updated_at FROM client"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.Client
	for rows.Next() {
		var c models.Client
		err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			continue
		}
		clients = append(clients, c)
	}
	return clients, nil
}

func (r *ClientRepository) FindByEmail(email string) (*models.Client, error) {
	query := "SELECT id, name, email, password_hash, created_at, updated_at FROM client WHERE email = ?"
	var c models.Client
	err := r.db.QueryRow(query, email).Scan(&c.ID, &c.Name, &c.Email, &c.PasswordHash, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}