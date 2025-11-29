package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moduloInicial/Aulas/18-UOW/internal/entity"
	"moduloInicial/Aulas/18-UOW/internal/repository"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// Mock implementations
type mockCourseRepo struct {
	insertCalled bool
	insertErr    error
	lastCourse   entity.Courses
}

func (m *mockCourseRepo) Insert(ctx context.Context, course entity.Courses) error {
	m.insertCalled = true
	m.lastCourse = course
	return m.insertErr
}

type mockCategoryRepo struct {
	insertCalled bool
	insertErr    error
	lastCategory entity.Category
}

func (m *mockCategoryRepo) Insert(ctx context.Context, category entity.Category) error {
	m.insertCalled = true
	m.lastCategory = category
	return m.insertErr
}

func TestExecute_Success(t *testing.T) {
	courseRepo := &mockCourseRepo{}
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

func TestExecute_CourseInsertError(t *testing.T) {
	courseRepo := &mockCourseRepo{insertErr: errors.New("course insert error")}
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

func TestExecute_CategoryInsertError(t *testing.T) {
	courseRepo := &mockCourseRepo{}
	categoryRepo := &mockCategoryRepo{insertErr: errors.New("category insert error")}
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

func TestExecute_Success_BD(t *testing.T) {

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
