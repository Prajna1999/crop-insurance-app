package repository

import (
	"github.com/Prajna1999/dataspace-be/internal/models"
	"gorm.io/gorm"
)

type EndpointRepository struct {
	DB *gorm.DB
}

func NewEndpointRepository(db *gorm.DB) *EndpointRepository {
	return &EndpointRepository{DB: db}
}

// add an endpoint to a specific api
func (r *EndpointRepository) Create(endpoint *models.EndPoint) error {
	return r.DB.Create(&endpoint).Error
}

// get details of a specific endpoint
func (r *EndpointRepository) GetEndpointByID(id uint) (*models.EndPoint, error) {
	var endpoint models.EndPoint
	err := r.DB.First(&endpoint, id).Error
	if err != nil {
		return nil, err
	}
	return &endpoint, nil
}
