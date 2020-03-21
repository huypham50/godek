package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetConnectionInfo -> return psqlInfo string after accessing env
func GetConnectionInfo() string {
	if err := godotenv.Load("dev.env"); err != nil {
		log.Fatal("Error loading .env files")
	}

	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	dbname := os.Getenv("PG_DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	return psqlInfo
}
