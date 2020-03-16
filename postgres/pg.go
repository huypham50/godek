package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// postgres driver dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InitDB -> Connect to local postgres database
func InitDB() {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	dbname := os.Getenv("PG_DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	// verify postgres and psqlInfo valid
	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.LogMode(true)
	log.Printf("connected to database %s on port %s", dbname, port)
}
