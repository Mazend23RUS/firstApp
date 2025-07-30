package models

import (
	"context"
	"fmt"
)

type Role struct {
	roleName    string
	permissions []string
}

func NewRole(
	ctx context.Context,
	roleName string, permission []string,
) (*Role, error) {

	if roleName != "Admin" && roleName != "User" || roleName == "" {
		return nil, fmt.Errorf("no valid role name")
	}

	ReadPermission := "read"
	WritePermission := "write"

	for _, c := range permission {
		if c != ReadPermission && c != WritePermission {
			return nil, fmt.Errorf("not valid permission")
		}
	}

	return &Role{
		roleName:    roleName,
		permissions: permission,
	}, nil
}

func (r *Role) RoleName() string {
	return r.roleName
}

func (r *Role) Permissions() []string {
	return r.permissions
}
