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

// ByID -> get deck by id
func (ds *DeckService) ByID(ID int) (*Deck, error) {
	var deck Deck

	err := ds.db.Where("id = ?", ID).First(&deck).Error

	switch err {
	case nil:
		return &deck, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// Create -> create provided deck
func (ds *DeckService) Create(deck *Deck) error {
	return ds.db.Create(deck).Error
}

// Update -> update deck
func (ds *DeckService) Update(deck *Deck) error {
	return ds.db.Save(deck).Error
}

// Delete -> delete requested deck
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
