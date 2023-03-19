package service

import (
	statsModel "github.com/linqcod/countries-telegram-bot/internal/userstatistics/model"
	"github.com/linqcod/countries-telegram-bot/internal/userstatistics/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetStatistics() (*statsModel.Statistics, error) {
	stats, err := s.repository.GetStatistics()
	if err != nil {
		return nil, err
	}

	return stats, nil
}
