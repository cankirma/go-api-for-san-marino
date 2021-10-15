package utils

import (
	"fmt"
	"github.com/cankirma/go-api-for-san-marino/pkg/role_repository"
)

func GetCredentialsByRole(role string) ([]string, error) {

	var credentials []string

	switch role {
	case role_repository.AdminRoleName:
		credentials = []string{
			role_repository.ProductCreateCredential,
			role_repository.ProductUpdateCredential,
			role_repository.ProductDeleteCredential,
			role_repository.CategoryCreateCredential,
			role_repository.CategoryUpdateCredential,
			role_repository.CategoryDeleteCredential,
		}
	case role_repository.ModeratorRoleName:
		credentials = []string{
			role_repository.ProductCreateCredential,
			role_repository.ProductUpdateCredential,
			role_repository.CategoryCreateCredential,
			role_repository.CategoryUpdateCredential,
		}
	case role_repository.UserRoleName:

		credentials = []string{
			role_repository.ProductCreateCredential,
			role_repository.CategoryCreateCredential,
		}
	default:
		// Return error message.
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}
