package models

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/phamstack/godek/lib/helpers"
)

var (
	// ErrNotFound -> no user in the db
	ErrNotFound = errors.New("models: resource not found")
)

// Auth -> return auth user and access token
// after a successful login
type Auth struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

// User -> schema-generated user struct
// will put into pg database as `users` table
type User struct {
	// gorm.model
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	// entities
	Email       string `json:"email" gorm:"not null;unique_index"`
	GoogleID    string `json:"googleId" gorm:"not null;unique_index"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Avatar      string `json:"avatar"`
	AccountType int    `json:"accountType"`
	// associations
	Decks []Deck `json:"decks"`
}

// UserService -> layer implementation of users
type UserService struct {
	db *gorm.DB
}

// NewUserService -> create a new instance of an UserService
// with error and db handling
func NewUserService(db *gorm.DB) UserService {
	return UserService{
		db: db,
	}
}

// ByGoogleID -> what might happen
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 - nil,otherError (something else went wrong -> 500 error)
func (us *UserService) ByGoogleID(googleID string) (*User, error) {
	var user User

	err := us.db.Where("google_id = ?", googleID).First(&user).Error

	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// ByEmail -> find user by email
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 - nil,otherError (something else went wrong -> 500 error)
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User

	err := us.db.Where("email = ?", email).First(&user).Error

	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// Create -> create provided user
func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

// Delete -> delete requested user
func (us *UserService) Delete(user *User) error {
	return us.db.Delete(user).Error
}

// Count -> total number of users
func (us *UserService) Count() int {
	var count int
	us.db.Model(&User{}).Count(&count)
	return count
}

// GenerateAuthToken -> generate access token when user logs in
func (us *UserService) GenerateAuthToken(user *User) string {
	if err := godotenv.Load("dev.env"); err != nil {
		log.Fatal("Error loading .env files")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.GoogleID,
		"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat": time.Now().Unix(),
	})
	tokenString, _ := token.SignedString([]byte(jwtSecret))

	helpers.LoggerLine(tokenString, user.Email)

	return tokenString
}

// ParseAuthToken -> decodes the token and returns user id
func (us *UserService) ParseAuthToken(jwtToken string) (string, error) {
	if err := godotenv.Load("dev.env"); err != nil {
		log.Fatal("Error loading .env files")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// type assertion -> else go specifies interface instead of string
		// https://stackoverflow.com/questions/14289256/cannot-convert-data-type-interface-to-type-string-need-type-assertion
		return claims["id"].(string), nil
	}

	return "", errors.New("Invalid JWT Token")
}

// Close -> closes the server database connection
func (us *UserService) Close() error {
	return us.db.Close()
}

// DestructiveReset -> drops the user table and rebuilds it
// dev only
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}
