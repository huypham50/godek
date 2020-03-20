package db

import "github.com/phamstack/godek/models"

// SeedDatabase -> Seed database for initial setup
func SeedDatabase(s *models.Services) {
	var users = []models.User{
		models.User{
			Email:    "peyton@manni.ng",
			Name:     "Peyton Manning",
			Username: "peyton0",
		},
		models.User{
			Email:    "drew@brees.no",
			Name:     "Drew Brees",
			Username: "drew1",
		},
		models.User{
			Email:    "jimmyg@49ers.com",
			Name:     "Jimmy Garoppolo",
			Username: "jimmy2",
		},
		models.User{
			Email:    "luck@blood.colts",
			Name:     "Andrew Luck",
			Username: "luck3",
		},
	}

	for _, user := range users {
		err := s.User.Create(&user)
		if err != nil {
			panic(err)
		}
	}
	return
}
