package repositories

import (
	"database/sql"
	"github.com/LaviqueDias/supermarket/internal/domain/models"
)

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(e *models.Employee) error {
	query := `INSERT INTO employe (name, email, password_hash, role) VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(query, e.Name, e.Email, e.PasswordHash, e.Role)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	e.ID = int(id)
	return nil
}

func (r *EmployeeRepository) FindAll() ([]models.Employee, error) {
	query := "SELECT id, name, email, role, created_at, updated_at FROM employe"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var e models.Employee
		err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Role, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			continue
		}
		employees = append(employees, e)
	}
	return employees, nil
}

func (r *EmployeeRepository) FindByEmail(email string) (*models.Employee, error) {
	query := "SELECT id, name, email, password_hash, role, created_at, updated_at FROM employe WHERE email = ?"
	var e models.Employee
	err := r.db.QueryRow(query, email).Scan(&e.ID, &e.Name, &e.Email, &e.PasswordHash, &e.Role, &e.CreatedAt, &e.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
