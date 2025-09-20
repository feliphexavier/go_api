package pkg

import (
	"database/sql"
	"fmt"
	"go_api/internal/config"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectPostgres(cfg *config.Config) (*sql.DB, error) {
	configString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	dbConnection, err := sql.Open("pgx", configString)
	if err != nil {
		log.Fatal(err)
	}
	err = dbConnection.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to database")
	return dbConnection, nil
}
