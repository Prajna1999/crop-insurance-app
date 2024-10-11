package repository

import (
	"github.com/Prajna1999/dataspace-be/internal/models"
	"gorm.io/gorm"
)

type OrganizationRepository struct {
	DB *gorm.DB
}

// inititalize the repo and write the queries
// takes a pointer to the database and spit out the repository

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{DB: db}
}

// create an organization
func (r *OrganizationRepository) Create(organization *models.Organization) error {
	return r.DB.Create(organization).Error
}

// get an organization by ID
func (r *OrganizationRepository) GetByID(id uint) (*models.Organization, error) {
	var organization models.Organization
	err := r.DB.First(&organization, id).Error

	if err != nil {
		return nil, err
	}
	return &organization, nil
}

// get all organizations
func (r *OrganizationRepository) GetAll() ([]*models.Organization, error) {
	var orgs []*models.Organization

	err := r.DB.Find(&orgs).Error

	if err != nil {
		return nil, err
	}
	// return a pointer to the orgs array
	return orgs, nil

}
