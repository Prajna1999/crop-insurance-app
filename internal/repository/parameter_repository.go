package repository

import (
	"github.com/Prajna1999/dataspace-be/internal/models"
	"gorm.io/gorm"
)

type ParameterRepository struct {
	DB *gorm.DB
}

func NewParameterRepository(db *gorm.DB) *ParameterRepository {
	return &ParameterRepository{DB: db}
}

// create parameter

func (r *ParameterRepository) Create(parameter *models.Parameter) error {
	return r.DB.Create(&parameter).Error
}

// get parameters by endpoint id
func (r *ParameterRepository) GetParametersByEndpointID(endpointID uint) ([]*models.Parameter, error) {
	var parameters []*models.Parameter

	err := r.DB.Where("end_point_id=?", endpointID).Find(&parameters).Error
	if err != nil {
		return nil, err
	}
	return parameters, nil
}

// update an existing parameter

// delete a parameter
