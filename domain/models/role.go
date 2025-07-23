package models

import (
	"context"
)

type Role struct {
	RoleName   string
	Permitions []string
}

func NewRole(
	ctx context.Context,
	roleName string, permition []string,
) (*Role, error) {

	return &Role{
		RoleName:   roleName,
		Permitions: permition,
	}, nil

}
