package models

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
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
	gorm.Model
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email" gorm:"not null;unique_index"`
	Avatar      string `json:"avatar"`
	AccountType int    `json:"accountType"`
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

// ByID -> what might happen
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 - nil,otherError (something else went wrong -> 500 error)
func (us *UserService) ByID(id uint) (*User, error) {
	var user User

	err := us.db.Where("id = ?", id).First(&user).Error

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

// Close -> closes the server database connection
func (us *UserService) Close() error {
	return us.db.Close()
}

// Demo -> closes the server database connection
func (us *UserService) Demo() error {
	fmt.Printf("%T\n", us.db.Create)
	return nil
}

// DestructiveReset -> drops the user table and rebuilds it
// dev only
func (us *UserService) DestructiveReset() {
	us.db.DropTableIfExists(&User{})
	us.db.AutoMigrate(&User{})
}
