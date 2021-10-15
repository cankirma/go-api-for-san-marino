package utils

import (
	"fmt"
	"github.com/cankirma/go-api-for-san-marino/pkg/role_repository"
)

func VerifyRole(role string) (string, error) {

	switch role {
	case role_repository.AdminRoleName:

	case role_repository.ModeratorRoleName:

	case role_repository.UserRoleName:

	default:

		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}
