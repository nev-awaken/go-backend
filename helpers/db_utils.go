package helpers

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func DbConnect(dbname string) (*sql.DB, error) {
	config, err := GetDbConfig()

	if err != nil {
		return nil, err
	}

	if dbname == "main" {
		dbname = config.DbName
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.PGUser, config.PGPassword, config.PGHost, config.PGPort, dbname,
	)

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

var dbPool *sql.DB

func InitializeConnectionPool(dbName string) error {
	conn, err := DbConnect(dbName)
	if err != nil {
		return err
	}
	conn.SetMaxOpenConns(20)
	conn.SetMaxIdleConns(20)
	conn.SetConnMaxLifetime(time.Hour)
	dbPool = conn
	return nil
}

func CheckoutConnection() *sql.DB {
	return dbPool
}
