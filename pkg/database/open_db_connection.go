package database

import "github.com/cankirma/go-api-for-san-marino/app/repositories"

type DbRepositories struct {
	*repositories.UserRepository
	*repositories.CategoryRepository
	*repositories.ProductRepository
}



func OpenDBConnection() (*DbRepositories, error) {
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &DbRepositories{

		UserRepository: &repositories.UserRepository{DB: db},
		CategoryRepository: &repositories.CategoryRepository{DB: db},
		ProductRepository: &repositories.ProductRepository{DB: db},
	}, nil
}
