package repository

import (
	"context"
	"database/sql"
	"moduloInicial/Aulas/18-UOW/internal/db"
	"moduloInicial/Aulas/18-UOW/internal/entity"
)

type CourseRepositoryInterface interface {
	Insert(ctx context.Context, course entity.Courses) error
}

type CourseRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCourseRepository(dtb *sql.DB) *CourseRepository {
	return &CourseRepository{
		DB:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *CourseRepository) Insert(ctx context.Context, course entity.Courses) error {
	return r.Queries.CreateCourse(ctx, db.CreateCourseParams{
		Name: course.Name,
	})
}
