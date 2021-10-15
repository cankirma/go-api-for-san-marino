package database

import (
	"fmt"
	"github.com/cankirma/go-api-for-san-marino/pkg/utils"
	"os"
	"strconv"
	"time"


	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v4/stdlib"
)


func PostgreSQLConnection() (*sqlx.DB, error) {

	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	postgresConnURL, err := utils.ConnectionURLBuilder("postgres")
	if err != nil {
		return nil, err
	}


	db, err := sqlx.Connect("pgx", postgresConnURL)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %v", err)
	}


	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn))


	if err := db.Ping(); err != nil {
		defer func(db *sqlx.DB) {
			err := db.Close()
			if err != nil {

			}
		}(db)
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
