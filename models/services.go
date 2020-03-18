package models

import (
	"github.com/jinzhu/gorm"

	// postgres dialect for gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Services -> single database connection to different services
type Services struct {
	User UserService
	db   *gorm.DB
}

// ServicesConfig -> configurations
type ServicesConfig func(*Services) error

// WithGorm -> init database connection
func WithGorm(dialect, connectionInfo string) ServicesConfig {
	return func(s *Services) error {
		db, err := gorm.Open(dialect, connectionInfo)
		if err != nil {
			return err
		}
		s.db = db
		return nil
	}
}

// WithLogMode -> specify the log mode
func WithLogMode(mode bool) ServicesConfig {
	return func(s *Services) error {
		s.db.LogMode(mode)
		return nil
	}
}

// WithUser -> init NewUserService
func WithUser() ServicesConfig {
	return func(s *Services) error {
		s.User = NewUserService(s.db)
		return nil
	}
}

// NewServices -> single source of truth
// umbrella db connection -> connect to other services
func NewServices(cfgs ...ServicesConfig) (*Services, error) {
	var s Services
	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}
	return &s, nil
}

// Close -> closes the database connection
func (s *Services) Close() error {
	return s.db.Close()
}

// DestructiveReset -> drops all tables and rebuilds them
func (s *Services) DestructiveReset() error {
	err := s.db.DropTableIfExists(&User{}).Error
	if err != nil {
		return err
	}
	return s.AutoMigrate()
}

// AutoMigrate -> reset database table
func (s *Services) AutoMigrate() error {
	return s.db.AutoMigrate(&User{}).Error
}
