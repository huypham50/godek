package db

import "github.com/phamstack/godek/models"

// SeedDatabase -> Seed database for initial setup
func SeedDatabase(s *models.Services) {
	var users = []models.User{
		models.User{
			GoogleID: "198923",
			Email:    "peyton@manni.ng",
			Name:     "Peyton Manning",
			Username: "peyton0",
		},
		models.User{
			GoogleID: "743823",
			Email:    "drew@brees.no",
			Name:     "Drew Brees",
			Username: "drew1",
		},
		models.User{
			GoogleID: "668323",
			Email:    "jimmyg@49ers.com",
			Name:     "Jimmy Garoppolo",
			Username: "jimmy2",
		},
		models.User{
			GoogleID: "007482",
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
