package main

import (
	"context"
	"database/sql"
	"fmt"
	"moduloInicial/Aulas/17-SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	dbconn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")

	if err != nil {
		panic(err.Error())

	}

	var queries = db.New(dbconn)

	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name:        "Backend",
		Description: sql.NullString{String: "cursinho de backend", Valid: true},
		ID:          uuid.NewString(),
	})

	if err != nil {
		panic(err.Error())
	}

	categories, err := queries.ListCategories(ctx)

	if err != nil {
		panic(err.Error())
	}

	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description.String)

	}

	for _, category := range categories {

		queries.UpdateCategories(ctx, db.UpdateCategoriesParams{
			ID:          category.ID,
			Name:        category.Name + " updated",
			Description: sql.NullString{String: category.Description.String + " updated", Valid: true},
		})
	}

	categories, _ = queries.ListCategories(ctx)

	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description.String)

	}

	for _, category := range categories {
		_ = queries.DeleteCategories(ctx, category.ID)
	}

	defer dbconn.Close()
}
