package service

import (
	"github.com/Prajna1999/dataspace-be/internal/models"
	"github.com/Prajna1999/dataspace-be/internal/repository"
)

type OrganizationService struct {
	repo *repository.OrganizationRepository
}

func NewOragnizationService(repo *repository.OrganizationRepository) *OrganizationService {
	return &OrganizationService{repo: repo}
}

// create organization by org_name, admin and string
func (s *OrganizationService) CreateOrganization(org_name string, admin_email string, admin_user_name string) error {
	organization := &models.Organization{
		OrgName:       org_name,
		AdminEmail:    admin_email,
		AdminUserName: admin_user_name,
	}
	return s.repo.Create(organization)
}

// get organization by org id
func (s *OrganizationService) GetOrganization(id uint) (*models.Organization, error) {
	return s.repo.GetByID(id)
}

// get all organizations
func (s *OrganizationService) GetAllOrganizations() ([]*models.Organization, error) {
	return s.repo.GetAll()
}
