package service

import (
	"context"

	"github.com/Frank-Macedo/13-GRPC/internal/database"
	"github.com/Frank-Macedo/13-GRPC/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDb database.Category
}

func NewCategoryService(categoryDb database.Category) *CategoryService {
	return &CategoryService{
		CategoryDb: categoryDb,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDb.Create(in.Name, in.Description)

	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDb.FindAll()
	if err != nil {
		return nil, err
	}

	var categoryResponse []*pb.Category

	for _, category := range categories {
		pbCategory := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
		categoryResponse = append(categoryResponse, pbCategory)
	}

	return &pb.CategoryList{
		Categories: categoryResponse,
	}, nil

}

func (c *CategoryService) SearchCategoryByCourseId(ctx context.Context, in *pb.SearchCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDb.FindByCourseId(in.CourseId)

	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}

func (c *CategoryService) SearchCategoryByName(ctx context.Context, in *pb.SearchCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDb.FindByName(in.Name)

	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}

func (c *CategoryService) GetCategoryById(ctx context.Context, in *pb.CategoryGetRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDb.FindById(in.CategoryId)

	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil
}

func (c *CategoryService) CreateCategorySream(stream pb.CategoryService_CreateCategorySreamServer) error {
	categories := &pb.CategoryList{}
	for {
		category, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				return stream.SendAndClose(categories)
			}
			return err
		}

		categoryResult, err := c.CategoryDb.Create(category.Name, category.Description)
		if err != nil {
			return err
		}
		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
	}
}

func (c *CategoryService) CreateCategoryStreamBothWays(stream pb.CategoryService_CreateCategoryStreamBothWaysServer) error {
	for {
		category, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				return nil
			}
			return err
		}

		categoryResult, err := c.CategoryDb.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.CategoryResponse{
			Category: &pb.Category{
				Id:          categoryResult.ID,
				Name:        categoryResult.Name,
				Description: categoryResult.Description,
			},
		})
		if err != nil {
			return err
		}
	}
}
