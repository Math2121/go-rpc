package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Courses struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryId  string
}

func (c *Courses) Create(name, description, categoryId string) (*Courses, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, categoryId)

	if err != nil {
		return nil, err
	}

	return &Courses{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryId:  categoryId,
	}, nil
}


func NewCourses(db *sql.DB) *Courses {
	return &Courses{db: db}
}
func (c *Courses) FindAll() ([]Courses, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses")

	if err != nil {
		return nil, err

	}

	defer rows.Close()

	courses := []Courses{}

	for rows.Next() {
		var id, name, description, categoryId string
		if err:= rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}
		courses = append(courses, Courses{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryId:  categoryId,
		})
	}
	return courses, nil
}

func (c *Courses) FindByCategoryID(categoryId string) ([]Courses, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = $1", categoryId)

	if err != nil {
		return nil, err

	}

	defer rows.Close()

	courses := []Courses{}

	for rows.Next(){
		var id, name, description, categoryId string

		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}
		courses = append(courses, Courses{
			ID:          id,
			Name:        name,
			Description: description,
			CategoryId:  categoryId,
		})
	}

	return courses, nil

}