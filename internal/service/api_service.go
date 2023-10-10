// internal/service/api_service.go
package service

import (
	"context"

	"ddvinyaninov/assets-tracker-api/internal/domain"
)

type apiService struct {
	apiRepository domain.ApiRepository
}

func NewApiService(apiRepository domain.ApiRepository) domain.ApiUsecase {
	return &apiService{apiRepository: apiRepository}
}

func (s apiService) List(ctx context.Context) ([]*domain.Api, error) {
	api, err := s.apiRepository.Select(ctx)
	if err != nil {
		return nil, err
	}

	return api, nil

}

func (s apiService) Create(ctx context.Context, body string) (*domain.Api, error) {
	api, err := s.apiRepository.Insert(ctx, body)
	if err != nil {
		return nil, err
	}

	return api, nil
}
