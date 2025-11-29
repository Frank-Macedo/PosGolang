package main

import (
	"context"
	"database/sql"
	"fmt"
	"moduloInicial/Aulas/17-SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, categoryArgs CategoryParams, courseArgs CourseParams) error {
	err := c.callTX(ctx, func(q *db.Queries) error {
		err := q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          categoryArgs.ID,
			Name:        categoryArgs.Name,
			Description: categoryArgs.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          courseArgs.ID,
			Name:        courseArgs.Name,
			Description: courseArgs.Description,
			CategoryID:  categoryArgs.ID,
			Price:       courseArgs.Price,
		})
		if err != nil {
			return err
		}
		return nil

	})

	if err != nil {
		return err
	}
	return nil
}

func (c *CourseDB) callTX(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)

	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %w, original error: %w", errRb, err)
		}
		return err
	}
	return tx.Commit()
}

func main() {
	ctx := context.Background()
	dbconn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")

	if err != nil {
		panic(err.Error())
	}

	coursedb := NewCourseDB(dbconn)

	defer coursedb.dbConn.Close()

	var categoryArgs = CategoryParams{
		ID:          uuid.NewString(),
		Name:        "Curso novo",
		Description: sql.NullString{"Mais um cursinho bolado", true},
	}

	var courseArgs = CourseParams{
		ID:          uuid.NewString(),
		Name:        "Curso novo",
		Description: sql.NullString{"Mais um cursinho bolado 2", true},
	}

	err = coursedb.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
	if err != nil {
		panic(err.Error())
	}

	listCourses, _ := coursedb.ListCourses(ctx)

	for _, v := range listCourses {
		fmt.Printf(v.ID, v.Name, v.Description, v.Price, v.CategoryID)
	}

}
