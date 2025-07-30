package models_test

import (
	"context"
	"testing"

	"github.com/alexey/firstApp/domain/models"
	"github.com/stretchr/testify/assert"
)

func TestCreatingNewRole(t *testing.T) {
	tests := []struct {
		Name         string
		RoleName     string
		Parmissions  []string
		wantErr      bool
		errorcontain string
	}{
		{
			Name:        "создание роли Admin",
			RoleName:    "Admin",
			Parmissions: []string{"read", "write"},
			wantErr:     false,
		},
		{
			Name:        "создание роли User",
			RoleName:    "User",
			Parmissions: []string{"read"},
			wantErr:     false,
		},
		{
			Name:         "невалидная роль",
			RoleName:     "invalid rolename",
			Parmissions:  []string{"read", "write"},
			wantErr:      true,
			errorcontain: "no valid role name",
		},
		{
			Name:         "невалидные допуски финциональности",
			RoleName:     "User",
			Parmissions:  []string{"update"},
			wantErr:      true,
			errorcontain: "not valid permission",
		},
	}

	for _, ts := range tests {

		t.Run(ts.Name, func(t *testing.T) {
			role, err := models.NewRole(context.Background(), ts.RoleName, ts.Parmissions)
			if ts.wantErr {
				assert.Error(t, err)
				if ts.errorcontain != "" {
					assert.Contains(t, err.Error(), ts.errorcontain)
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, ts.RoleName, role.RoleName())
			assert.Equal(t, ts.Parmissions, role.Permissions())
		})

	}

}
