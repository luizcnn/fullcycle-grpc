package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

func (category *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := category.db.Exec(
		"INSERT INTO CATEGORIES (id, name, description) values ($1, $2, $3)", id, name, description,
	)

	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM CATEGORIES")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return categories, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}

	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
	var id, name, description string
	err := c.db.QueryRow("SELECT ca.id, ca.name, ca.description FROM CATEGORIES ca INNER JOIN COURSES co ON ca.id = co.category_id where co.id = $1", courseID).
		Scan(&id, &name, &description)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}
