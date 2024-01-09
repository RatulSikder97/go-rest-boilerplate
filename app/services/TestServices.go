package services

import (
	"com/beeda/search/app/models"
	"com/beeda/search/app/repositories"
)

type TestService struct {
	Repository *repositories.TestRepository
}

func NewTestService(repo *repositories.TestRepository) *TestService {
	return &TestService{Repository: repo}
}

// GetAll returns all records from the model using the repository
func (s *TestService) GetAll() ([]models.TestModel, error) {
	return s.Repository.GetAll()
}
