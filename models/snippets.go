package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Snippet -> notes
// will put into pg database as `snippets` table
type Snippet struct {
	// gorm.Model
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	UserID int `json:"userId" gorm:"index:user_id_snippet;not null"`
	DeckID int `json:"deckId" gorm:"index:deck_id_snippet"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Archive     bool   `json:"archive"`
}

// SnippetService -> layer implementation of snippets
type SnippetService struct {
	db *gorm.DB
}

// NewSnippetService -> create a new &unique instance of an SnippetService
func NewSnippetService(db *gorm.DB) SnippetService {
	return SnippetService{
		db: db,
	}
}

// ByID -> get deck by id
func (ss *SnippetService) ByID(ID int) (*Snippet, error) {
	var snippet Snippet

	err := ss.db.Where("id = ?", ID).First(&snippet).Error

	switch err {
	case nil:
		return &snippet, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// Create -> create provided deck
func (ss *SnippetService) Create(snippet *Snippet) error {
	return ss.db.Create(snippet).Error
}

// Update -> update deck
func (ss *SnippetService) Update(snippet *Snippet) error {
	return ss.db.Save(snippet).Error
}

// Delete -> delete requested deck
func (ss *SnippetService) Delete(snippet *Snippet) error {
	return ss.db.Delete(snippet).Error
}

// Filter -> search decks from user id
func (ss *SnippetService) Filter(userID int) ([]*Snippet, error) {
	var snippets []*Snippet

	// deadlinest first desc vs asc
	if err := ss.db.Where("user_id = ?", userID).Order("updated_at asc").Find(&snippets).Error; err != nil {
		return nil, err
	}

	return snippets, nil
}
