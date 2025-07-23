package repository

import (
	"context"

	"github.com/alexey/firstApp/domain/models"
)

type InterfaceUserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	Save(ctx context.Context, user *models.User) error
}
