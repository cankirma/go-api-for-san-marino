package repositories

import (
	"github.com/cankirma/go-api-for-san-marino/app/entities"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)


type UserRepository struct {
	*sqlx.DB
}


func (q *UserRepository) GetUserByID(id uuid.UUID) (entities.User, error) {

	user := entities.User{}


	query := `SELECT * FROM users WHERE id = $1`


	err := q.Get(&user, query, id)
	if err != nil {

		return user, err
	}

	return user, nil
}

func (q *UserRepository) GetUserByEmail(email string) (entities.User, error) {

	user := entities.User{}


	query := `SELECT * FROM users WHERE email = $1`


	err := q.Get(&user, query, email)
	if err != nil {

		return user, err
	}


	return user, nil
}


func (q *UserRepository) CreateUser(u *entities.User) error {

	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := q.Exec(
		query,
		u.ID, u.CreatedAt, u.UpdatedAt, u.Email, u.PasswordHash, u.UserStatus, u.UserRole,
	)
	if err != nil {

		return err
	}


	return nil
}
