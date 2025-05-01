package setup

import (
	"context"
	"errors"
	"log"
	"strings"
	"web_app_fiber/helpers"

	"github.com/jackc/pgconn"
)

func CreateDb() {
	dbName := "task_manager"
	query := "CREATE DATABASE " + dbName

	conn, err := helpers.DbConnect("postgres")
	if err != nil {
		log.Fatal("Failed to connect to Postgres: ", err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(context.Background(), query)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "42P04" {
				log.Printf("Database %s already exists. Skipping creation.", dbName)
				return
			}
		}
		if err.Error() != "" && (containsIgnoreCase(err.Error(), "already exists") || containsIgnoreCase(err.Error(), "42P04")) {
			log.Printf("Database %s already exists (caught by string parsing). Skipping creation.", dbName)
			return
		}
		log.Fatal("Failed to create database: ", err)
	}
}

func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func CreateTables() {
	conn, err := helpers.DbConnect("main")
	if err != nil {
		log.Fatal("Failed to connect to task_manager DB: ", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		completed BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT NOW()
	);`

	_, err = conn.ExecContext(context.Background(), query)
	defer conn.Close()
	if err != nil {
		log.Fatal("Failed to create tables: ", err)
	}

	log.Println("Tables created successfully (if not existing).")
}
