package repository

import (
	"context"

	"github.com/alexey/firstApp/domain/models"
)

type RepositoryOfPermission interface {
	GetPermissions(ctx context.Context, nameRole string) ([]models.Role, error)
}
