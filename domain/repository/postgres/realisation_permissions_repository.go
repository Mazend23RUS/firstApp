package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alexey/firstApp/domain/models"
	"github.com/alexey/firstApp/domain/repository"
)

type PermissionRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) repository.RepositoryOfPermission {
	return &PermissionRepository{db: db}
}

func (r *PermissionRepository) GetPermissions(ctx context.Context, roleName string) ([]models.Role, error) {
	query := `SELECT permission FROM role_permissions WHERE role_name = $1`

	rows, err := r.db.QueryContext(ctx, query, roleName)
	if err != nil {
		return nil, fmt.Errorf("failed to query permissions: %w", err)
	}
	defer rows.Close()

	var permissions []string
	for rows.Next() {
		var permission string
		if err := rows.Scan(&permission); err != nil {
			return nil, fmt.Errorf("failed to scan permission: %w", err)
		}
		permissions = append(permissions, permission)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("В получении строк ошибка : %w", err)
	}

	role, err := models.NewRole(ctx, roleName, permissions)
	if err != nil {
		return nil, fmt.Errorf("Ошибка в создании роли:  %w", err)

	}

	return []models.Role{*role}, nil
}
