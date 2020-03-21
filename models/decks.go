package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Deck -> labels/tags
// will put into pg database as `decks` table
type Deck struct {
	// gorm.Model
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	UserID int `json:"userId" gorm:"index:user_id;not null"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Label       string `json:"label" gorm:"not null"`
	Color       string `json:"color"`
	Archive     bool   `json:"archive"`
}

// DeckService -> layer implementation of decks
type DeckService struct {
	db *gorm.DB
}

// NewDeckService -> create a new &unique instance of an DeckService
func NewDeckService(db *gorm.DB) DeckService {
	return DeckService{
		db: db,
	}
}

// Create -> create provided user
func (ds *DeckService) Create(deck *Deck) error {
	return ds.db.Create(deck).Error
}

// Delete -> delete requested user
func (ds *DeckService) Delete(deck *Deck) error {
	return ds.db.Delete(deck).Error
}

// Filter -> search decks from user id
func (ds *DeckService) Filter(userID int) ([]*Deck, error) {
	var decks []*Deck
	if err := ds.db.Where("user_id = ?", userID).Find(&decks).Error; err != nil {
		return nil, err
	}

	return decks, nil
}
