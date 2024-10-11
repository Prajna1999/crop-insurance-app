package service

import (
	"github.com/Prajna1999/dataspace-be/internal/models"
	"github.com/Prajna1999/dataspace-be/internal/repository"
)

type EndpointService struct {
	endpointRepo *repository.EndpointRepository
	paramRepo    *repository.ParameterRepository
}

func NewEndpointService(endpointRepo *repository.EndpointRepository, paramRepo *repository.ParameterRepository) *EndpointService {
	return &EndpointService{endpointRepo: endpointRepo, paramRepo: paramRepo}
}

func (s *EndpointService) CreateEndpoint(endpoint *models.EndPoint) error {
	return s.endpointRepo.Create(endpoint)
}

func (s *EndpointService) GetEndpointByID(id uint) (*models.EndPoint, error) {
	return s.endpointRepo.GetEndpointByID(id)
}

// add a parameter to an endpoint
func (s *EndpointService) AddParameterToEndpoint(endpointID uint, parameter *models.Parameter) error {
	// get the endpoint by ID
	endPoint, err := s.endpointRepo.GetEndpointByID(endpointID)
	if err != nil {
		return err
	}
	parameter.EndPointID = endPoint.ID
	return s.paramRepo.Create(parameter)

}

// List all parameter for a specific endpoint
func (s *EndpointService) GetAllParametersForEndpoint(endpointID uint) ([]*models.Parameter, error) {
	return s.paramRepo.GetParametersByEndpointID(endpointID)
}
