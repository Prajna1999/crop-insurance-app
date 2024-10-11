package repository

import (
	"github.com/Prajna1999/dataspace-be/internal/models"
	"gorm.io/gorm"
)

type ApiRepository struct {
	DB *gorm.DB
}

func NewApiRepository(db *gorm.DB) *ApiRepository {
	return &ApiRepository{DB: db}
}

func (r *ApiRepository) Create(api *models.Api) error {
	return r.DB.Create(api).Error
}

func (r *ApiRepository) GetAll(filters map[string]interface{}) ([]*models.Api, error) {
	var apis []*models.Api
	query := r.DB

	// enumerate all the filters and add them to the query
	for key, value := range filters {
		query = query.Where(key, value)

	}
	err := query.Find(&apis).Error
	if err != nil {
		return nil, err
	}
	return apis, nil
}

func (r *ApiRepository) GetByID(id uint) (*models.Api, error) {
	var api models.Api
	err := r.DB.First(&api, id).Error

	if err != nil {
		return nil, err
	}
	return &api, nil

}
