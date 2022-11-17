package service

import (
	"context"
	"inventory/internal/models"
)

type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
	
}
