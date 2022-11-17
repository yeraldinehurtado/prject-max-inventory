package service

import (
	"context"
	"inventory/internal/models"
	"inventory/internal/repository"
)

// service is the business logic of the application
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
        repo: repo,
    }
}
