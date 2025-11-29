package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moduloInicial/Aulas/18-UOW/internal/db"
	"moduloInicial/Aulas/18-UOW/internal/entity"
	"moduloInicial/Aulas/18-UOW/internal/repository"
	"moduloInicial/Aulas/18-UOW/sql/pkg/uow"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// Mock implementations
type mockCourseRepoUow struct {
	insertCalled bool
	insertErr    error
	lastCourse   entity.Courses
}

func (m *mockCourseRepoUow) Insert(ctx context.Context, course entity.Courses) error {
	m.insertCalled = true
	m.lastCourse = course
	return m.insertErr
}

type mockCategoryRepoUow struct {
	insertCalled bool
	insertErr    error
	lastCategory entity.Category
}

func (m *mockCategoryRepoUow) Insert(ctx context.Context, category entity.Category) error {
	m.insertCalled = true
	m.lastCategory = category
	return m.insertErr
}

func TestAddCourseUow(t *testing.T) {

	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")

	dbt.Exec("DROP TABLE IF EXISTS `courses`;")
	dbt.Exec("DROP TABLE IF EXISTS `categories`;")

	dbt.Exec(`CREATE TABLE categories (
		id INT AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id)
	);`)

	dbt.Exec(`CREATE TABLE courses (
		id INT AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		category_id INT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (id),
		FOREIGN KEY (category_id) REFERENCES categories(id)
	);`)
	if err != nil {
		fmt.Println("erro ao abrir conexão com sql")
		t.FailNow()
	}

	ctx := context.Background()
	uow := uow.NewUow(ctx, nil)

	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}

	uow.Register("CourseRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(dbt)
		repo.Queries = db.New(tx)
		return repo

	})
	uow.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})

	useCase := NewAddCourseUseCaseUow(uow)

	input := InputUseCaseUow{
		CategoryName:     "Programming",
		CourseName:       "Go Basics",
		CourseCategoryId: 1,
	}
	err = useCase.ExecuteUow(ctx, input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

}
func TestExecute_Success_Uow(t *testing.T) {
	courseRepo := &mockCourseRepoUow{}
	categoryRepo := &mockCategoryRepo{}
	useCase := NewAddCourseUseCase(courseRepo, categoryRepo)

	input := InputUseCase{
		CategoryName:     "Programming",
		CourseName:       "Go Basics",
		CourseCategoryId: 1,
	}

	err := useCase.Execute(context.Background(), input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !courseRepo.insertCalled {
		t.Error("expected CourseRepository.Insert to be called")
	}
	if !categoryRepo.insertCalled {
		t.Error("expected CategoryRepository.Insert to be called")
	}
	if courseRepo.lastCourse.Name != input.CourseName {
		t.Errorf("expected course name %q, got %q", input.CourseName, courseRepo.lastCourse.Name)
	}
	if categoryRepo.lastCategory.Name != input.CategoryName {
		t.Errorf("expected category name %q, got %q", input.CategoryName, categoryRepo.lastCategory.Name)
	}
	if categoryRepo.lastCategory.ID != input.CourseCategoryId {
		t.Errorf("expected category ID %d, got %d", input.CourseCategoryId, categoryRepo.lastCategory.ID)
	}
}

func TestExecute_CourseInsertError_Uow(t *testing.T) {
	courseRepo := &mockCourseRepoUow{insertErr: errors.New("course insert error")}
	categoryRepo := &mockCategoryRepo{}
	useCase := NewAddCourseUseCase(courseRepo, categoryRepo)

	input := InputUseCase{
		CategoryName:     "Programming",
		CourseName:       "Go Basics",
		CourseCategoryId: 1,
	}

	err := useCase.Execute(context.Background(), input)
	if err == nil || err.Error() != "course insert error" {
		t.Fatalf("expected course insert error, got %v", err)
	}
	if !courseRepo.insertCalled {
		t.Error("expected CourseRepository.Insert to be called")
	}
	if categoryRepo.insertCalled {
		t.Error("expected CategoryRepository.Insert NOT to be called")
	}
}

func TestExecute_CategoryInsertError_Uow(t *testing.T) {
	courseRepo := &mockCourseRepoUow{}
	categoryRepo := &mockCategoryRepoUow{insertErr: errors.New("category insert error")}
	useCase := NewAddCourseUseCase(courseRepo, categoryRepo)

	input := InputUseCase{
		CategoryName:     "Programming",
		CourseName:       "Go Basics",
		CourseCategoryId: 1,
	}

	err := useCase.Execute(context.Background(), input)
	if err == nil || err.Error() != "category insert error" {
		t.Fatalf("expected category insert error, got %v", err)
	}
	if !courseRepo.insertCalled {
		t.Error("expected CourseRepository.Insert to be called")
	}
	if !categoryRepo.insertCalled {
		t.Error("expected CategoryRepository.Insert to be called")
	}
}

func TestExecute_Success_BD_Uow(t *testing.T) {

	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")

	if err != nil {
		fmt.Println("erro ao abrir conexão com sql")
		t.FailNow()
	}

	courseRepo := repository.NewCourseRepository(dbt)
	categoryRepo := repository.NewCategoryRepository(dbt)

	useCase := NewAddCourseUseCase(courseRepo, categoryRepo)

	input := InputUseCase{
		CategoryName:     "Programming",
		CourseName:       "Go Basics",
		CourseCategoryId: 1,
	}

	err = useCase.Execute(context.Background(), input)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
