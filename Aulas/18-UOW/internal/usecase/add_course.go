package usecase

import (
	"context"
	"moduloInicial/Aulas/18-UOW/internal/entity"
	"moduloInicial/Aulas/18-UOW/internal/repository"
)

type InputUseCase struct {
	CategoryName     string
	CourseName       string
	CourseCategoryId int
}

type AddCourseUseCase struct {
	CourseRepository   repository.CourseRepositoryInterface
	CategoryRepository repository.CategoryRepositoryInterface
}

func NewAddCourseUseCase(courseRepository repository.CourseRepositoryInterface, categoryRepository repository.CategoryRepositoryInterface) *AddCourseUseCase {
	return &AddCourseUseCase{
		CourseRepository:   courseRepository,
		CategoryRepository: categoryRepository,
	}
}

func (ad *AddCourseUseCase) Execute(ctx context.Context, input InputUseCase) error {

	course := entity.Courses{
		Name: input.CourseName,
	}

	category := entity.Category{
		ID:   input.CourseCategoryId,
		Name: input.CategoryName,
	}

	err := ad.CategoryRepository.Insert(ctx, category)
	if err != nil {
		return err
	}

	err = ad.CourseRepository.Insert(ctx, course)
	if err != nil {
		return err
	}

	return nil
}
