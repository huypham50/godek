package models

import (
	"net/http"
	"time"

	readability "github.com/go-shiori/go-readability"
	"github.com/jinzhu/gorm"
)

// Bookmark -> saved articles
// will put into pg database as `bookmarks` table
type Bookmark struct {
	// gorm.Model
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	UserID int `json:"userId" gorm:"index:user_id_bookmark;not null"`
	DeckID int `json:"deckId" gorm:"index:deck_id_bookmark"`

	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	WordCount   int    `json:"wordCount"`
	Archive     bool   `json:"archive"`
}

// BookmarkService -> layer implementation of decks
type BookmarkService struct {
	db *gorm.DB
}

// NewBookmarkService -> create a new &unique instance of an BookmarkService
func NewBookmarkService(db *gorm.DB) BookmarkService {
	return BookmarkService{
		db: db,
	}
}

// ByID -> get deck by id
func (bs *BookmarkService) ByID(ID int) (*Bookmark, error) {
	var bookmark Bookmark

	err := bs.db.Where("id = ?", ID).First(&bookmark).Error

	switch err {
	case nil:
		return &bookmark, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FetchURL -> fetch metadata first and give it back to user
func (bs *BookmarkService) FetchURL(url string) (*Bookmark, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	article, err := readability.FromReader(res.Body, url)
	if err != nil {
		return nil, err
	}

	fetchedBookmark := &Bookmark{
		URL:         url,
		Title:       article.Title,
		Description: article.Excerpt,
		Thumbnail:   article.Image,
		WordCount:   article.Length,
	}

	return fetchedBookmark, nil
}

// Create -> create provided deck
func (bs *BookmarkService) Create(bookmark *Bookmark) error {
	return bs.db.Create(bookmark).Error
}

// Update -> update deck
func (bs *BookmarkService) Update(bookmark *Bookmark) error {
	return bs.db.Save(bookmark).Error
}

// Delete -> delete requested deck
func (bs *BookmarkService) Delete(bookmark *Bookmark) error {
	return bs.db.Delete(bookmark).Error
}

// Filter -> search decks from user id
func (bs *BookmarkService) Filter(userID int) ([]*Bookmark, error) {
	var bookmarks []*Bookmark

	// newest first! desc vs asc
	if err := bs.db.Where("user_id = ?", userID).Order("created_at desc").Find(&bookmarks).Error; err != nil {
		return nil, err
	}

	return bookmarks, nil
}
