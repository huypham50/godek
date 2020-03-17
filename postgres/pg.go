package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/phamstack/godek/models"

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
	us, err := models.NewUserService(psqlInfo)

	if err != nil {
		panic(err)
	}
	defer us.Close()

	us.DestructiveReset()
	log.Printf("connected to database %s on port %s", dbname, port)

	tom := &models.User{
		Email:    "tom@brady.ne",
		Name:     "Thomas Brady",
		Username: "cheat",
	}
	if err := us.Create(tom); err != nil {
		panic(err)
	}
}
