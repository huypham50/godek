package models

import "time"

// Deck -> labels/tags
// will put into pg database as `decks` table
type Deck struct {
	// gorm.Model
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	UserID string `json:userId gorm:"not null;unique_index"`

	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Label   string `json:"label"`
	Color   string `json:"color"`
	Archive bool   `json:"archive"`
}
