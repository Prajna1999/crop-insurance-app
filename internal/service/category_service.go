package service

import (
	"github.com/Prajna1999/dataspace-be/internal/models"
	"github.com/Prajna1999/dataspace-be/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// create category service by category name
func (s *CategoryService) CreateCategory(name string, description string) error {
	category := &models.Category{
		Name:        name,
		Description: description,
	}

	return s.repo.Create(category)
}

// get all categories
func (s *CategoryService) GetAllCategories() ([]*models.Category, error) {
	return s.repo.GetAllCategories()
}

//get one category by id

func (s *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	return s.repo.GetCategoryByID(id)
}
