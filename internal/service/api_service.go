package service

import (
	"github.com/Prajna1999/dataspace-be/internal/models"
	"github.com/Prajna1999/dataspace-be/internal/repository"
)

type ApiService struct {
	apiRepo      *repository.ApiRepository
	endpointRepo *repository.EndpointRepository
}

func NewApiService(apiRepo *repository.ApiRepository, endpointRepo *repository.EndpointRepository) *ApiService {
	return &ApiService{apiRepo: apiRepo, endpointRepo: endpointRepo}
}

// create a new api
func (s *ApiService) CreateApi(api *models.Api) error {
	return s.apiRepo.Create(api)
}

// get all apis or filtered
func (s *ApiService) GetAllApis(filters map[string]interface{}) ([]*models.Api, error) {
	return s.apiRepo.GetAll(filters)
}

// add new endpoint to an existing api
func (s *ApiService) AddEndpointToApi(apiID uint, endpoint *models.EndPoint) error {
	api, err := s.apiRepo.GetByID(apiID)

	if err != nil {
		return err
	}
	endpoint.ApiID = api.ID

	return s.endpointRepo.Create(endpoint)
}
