package service

import (
	"github.com/Prajna1999/dataspace-be/internal/models"
	"github.com/Prajna1999/dataspace-be/internal/repository"
)

type ParameterService struct {
	paramRepo *repository.ParameterRepository
}

func NewParameterService(repo *repository.ParameterRepository) *ParameterService {
	return &ParameterService{paramRepo: repo}
}

func (s *ParameterService) CreateParameter(param *models.Parameter) error {
	return s.paramRepo.Create(param)
}
