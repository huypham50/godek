package db

import (
	"time"

	"github.com/phamstack/godek/models"
)

// SeedDatabase -> Seed database for initial setup
func SeedDatabase(s *models.Services) {
	var users = []models.User{
		models.User{
			GoogleID: "198923",
			Email:    "peyton@manni.ng",
			Name:     "Peyton Manning",
		},
		models.User{
			GoogleID: "743823",
			Email:    "drew@brees.no",
			Name:     "Drew Brees",
		},
		models.User{
			GoogleID: "668323",
			Email:    "jimmyg@49ers.com",
			Name:     "Jimmy Garoppolo",
		},
		models.User{
			GoogleID: "007482",
			Email:    "luck@blood.colts",
			Name:     "Andrew Luck",
		},
		models.User{
			GoogleID: "111172427755155985046",
			Email:    "cocwedc@gmail.com",
			Name:     "Pham Huy",
		},
	}

	var decks = []models.Deck{
		models.Deck{
			UserID:      1,
			Title:       "Indianapolis Colts",
			Description: "Carrying The Colts",
			Label:       "IND",
			Color:       "#ece",
		},
		models.Deck{
			UserID:      1,
			Title:       "Denver Broncos",
			Description: "Defense Win Championships",
			Label:       "DEN",
			Color:       "#e67",
		},
		models.Deck{
			UserID:      3,
			Title:       "New England Patriots",
			Description: "Being Tom Brady's Sub",
			Label:       "NE",
			Color:       "#c11",
		},
		models.Deck{
			UserID:      3,
			Title:       "San Francisco 49ers",
			Description: "Riding the Defense Till' The End",
			Label:       "SF",
			Color:       "#444",
		},
		models.Deck{
			UserID:      5,
			Title:       "Baltimore Ravens",
			Description: "New Era",
			Label:       "BAL",
			Color:       "purple",
		},
		models.Deck{
			UserID:      5,
			Title:       "Seattle Seahawks",
			Description: "Run The Football",
			Label:       "SEATLLE",
			Color:       "green",
		},
		models.Deck{
			UserID:      5,
			Title:       "Denver Broncos",
			Description: "Homeboy Done Good",
			Label:       "DEN",
			Color:       "orange",
		},
	}

	var todos = []models.Todo{
		models.Todo{
			UserID:      5,
			DeckID:      6,
			Title:       "Eating Lunch",
			Description: "12pm",
			Deadline:    time.Now().Add(time.Hour * 100),
		},
		models.Todo{
			UserID:      5,
			DeckID:      6,
			Title:       "Having Fun",
			Description: "all day",
			Deadline:    time.Now().Add(time.Hour * 15),
		},
		models.Todo{
			UserID:      5,
			DeckID:      7,
			Title:       "Practice Bunch",
			Description: "bruh",
			Deadline:    time.Now().Add(time.Hour * 48),
		},
	}

	var bookmarks = []models.Bookmark{
		models.Bookmark{
			UserID:      3,
			Title:       "Demot",
			Description: "Demod",
			WordCount:   7575,
		},
	}

	for _, user := range users {
		if err := s.User.Create(&user); err != nil {
			panic(err)
		}
	}

	for _, deck := range decks {
		if err := s.Deck.Create(&deck); err != nil {
			panic(err)
		}
	}

	for _, todo := range todos {
		if err := s.Todo.Create(&todo); err != nil {
			panic(err)
		}
	}

	for _, bookmark := range bookmarks {
		if err := s.Bookmark.Create(&bookmark); err != nil {
			panic(err)
		}
	}

	return
}
