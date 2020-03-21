package db

import "github.com/phamstack/godek/models"

// SeedDatabase -> Seed database for initial setup
func SeedDatabase(s *models.Services) {
	var users = []models.User{
		models.User{
			GoogleID: "198923",
			Email:    "peyton@manni.ng",
			Name:     "Peyton Manning",
			Username: "peyton1",
		},
		models.User{
			GoogleID: "743823",
			Email:    "drew@brees.no",
			Name:     "Drew Brees",
			Username: "drew2",
		},
		models.User{
			GoogleID: "668323",
			Email:    "jimmyg@49ers.com",
			Name:     "Jimmy Garoppolo",
			Username: "jimmy3",
		},
		models.User{
			GoogleID: "007482",
			Email:    "luck@blood.colts",
			Name:     "Andrew Luck",
			Username: "luck4",
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
			UserID:      3,
			Title:       "New England Patriots",
			Description: "Second stint",
			Label:       "NE2",
			Color:       "#c11",
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
	return
}
