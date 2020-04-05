package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Todo -> todo list
// will put into pg database as `todos` table
type Todo struct {
	// gorm.Model
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	UserID int `json:"userId" gorm:"index:user_id_todo;not null"`
	DeckID int `json:"deckId" gorm:"index:deck_id_todo"`

	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Complete    bool      `json:"archive"`
}

// TodoService -> layer implementation of decks
type TodoService struct {
	db *gorm.DB
}

// NewTodoService -> create a new &unique instance of an TodoService
func NewTodoService(db *gorm.DB) TodoService {
	return TodoService{
		db: db,
	}
}

// ByID -> get deck by id
func (ts *TodoService) ByID(ID int) (*Todo, error) {
	var todo Todo

	err := ts.db.Where("id = ?", ID).First(&todo).Error

	switch err {
	case nil:
		return &todo, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// Create -> create provided deck
func (ts *TodoService) Create(todo *Todo) error {
	return ts.db.Create(todo).Error
}

// Update -> update deck
func (ts *TodoService) Update(todo *Todo) error {
	return ts.db.Save(todo).Error
}

// Delete -> delete requested deck
func (ts *TodoService) Delete(todo *Todo) error {
	return ts.db.Delete(todo).Error
}

// Filter -> search decks from user id
func (ts *TodoService) Filter(userID int) ([]*Todo, error) {
	var todos []*Todo

	// deadlinest first desc vs asc
	if err := ts.db.Where("user_id = ?", userID).Order("deadline asc").Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}
