package handler

import (
	"github.com/madsilver/api-challenge/internal/infra/repository"
	"github.com/madsilver/api-challenge/internal/usecase"
)

type Manager struct {
	MetricHandler MetricHandler
	UserHandler   UserHandler
}

func NewManager() *Manager {
	metricHandler := NewMetricHandler()
	userHandler := NewUserHandler(
		usecase.NewUserUseCase(
			repository.NewUserRepository(),
		),
	)

	return &Manager{
		MetricHandler: metricHandler,
		UserHandler:   userHandler,
	}
}
