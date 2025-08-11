package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	Id          string
	Name        string
	Description string
	CategoryId  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name string, description string, categoryId string) (*Course, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("insert into Courses (Id, Name, Description, CategoryId) values ($1, $2 ,$3, $4)", id, name, description, categoryId)
	if err != nil {
		return nil, err
	}

	return &Course{Id: id, Name: name, Description: description, CategoryId: categoryId}, nil

}

func (c *Course) FindAll() ([]Course, error) {

	rows, err := c.db.Query("select * from courses")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	courses := []Course{}

	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		courses = append(courses, Course{Id: id, Name: name, Description: description, CategoryId: categoryID})
	}
	return courses, nil

}

func (c *Course) FindByCategoryId(categoryId string) ([]Course, error) {

	rows, err := c.db.Query("select * from courses where categoryId = $1", categoryId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	courses := []Course{}

	for rows.Next() {
		var id, name, description, categoryID string
		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		courses = append(courses, Course{Id: id, Name: name, Description: description, CategoryId: categoryID})
	}
	return courses, nil

}
