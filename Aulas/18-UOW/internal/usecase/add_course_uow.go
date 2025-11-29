package usecase

import (
	"context"
	"moduloInicial/Aulas/18-UOW/internal/entity"
	"moduloInicial/Aulas/18-UOW/internal/repository"
	"moduloInicial/Aulas/18-UOW/sql/pkg/uow"
)

type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryId int
}

type AddCourseUseCaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseUseCaseUow(uow uow.UowInterface) *AddCourseUseCaseUow {
	return &AddCourseUseCaseUow{
		Uow: uow,
	}
}

func (ad *AddCourseUseCaseUow) ExecuteUow(ctx context.Context, input InputUseCaseUow) error {

	return ad.Uow.Do(ctx, func(uow *uow.Uow) error {

		repoCategory := ad.getCategoryRepository(ctx)
		repoCourse := ad.getCourseRepository(ctx)

		category := entity.Category{
			Name: input.CategoryName,
		}

		course := entity.Courses{
			Name: input.CourseName,
		}

		err := repoCategory.Insert(ctx, category)
		if err != nil {
			return err
		}

		err = repoCourse.Insert(ctx, course)
		if err != nil {
			return err
		}
		return nil
	})

	// err := ad.CategoryRepository.Insert(ctx, category)
	// if err != nil {
	// 	return err
	// }

	// err = ad.CourseRepository.Insert(ctx, course)
	// if err != nil {
	// 	return err
	// }

	// return nil
}

func (ad *AddCourseUseCaseUow) getCategoryRepository(ctx context.Context) repository.CategoryRepository {
	repo, err := ad.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err.Error())
	}

	return repo.(repository.CategoryRepository)
}

func (ad *AddCourseUseCaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := ad.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err.Error())
	}

	return repo.(repository.CourseRepositoryInterface)
}
