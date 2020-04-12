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
			UserID: 1,
			Label:  "IND",
			Color:  "#ece",
		},
		models.Deck{
			UserID: 1,
			Label:  "DEN",
			Color:  "#e67",
		},
		models.Deck{
			UserID: 3,
			Label:  "NE",
			Color:  "#c11",
		},
		models.Deck{
			UserID: 3,
			Label:  "SF",
			Color:  "#444",
		},
		models.Deck{
			UserID: 5,
			Label:  "BAL",
			Color:  "purple",
		},
		models.Deck{
			UserID: 5,
			Label:  "SEATLLE",
			Color:  "green",
		},
		models.Deck{
			UserID: 5,
			Label:  "DEN",
			Color:  "orange",
		},
	}

	var snippets = []models.Snippet{
		models.Snippet{
			UserID:      5,
			DeckID:      7,
			Title:       "Shopping List For This Week",
			Description: `[{"children":[{"text":"Sit ut fugit et quae. Velit rem esse molestias rerum. Animi ut animi. Molestias occaecati dolores sed repudiandae mollitia. Officiis maxime a dolores quas."}]},{"children":[{"text":""}]},{"children":[{"text":"Quisquam accusantium voluptates quia officia accusamus nihil voluptas perspiciatis repellendus. Dolor sit adipisci minima amet hic consequatur explicabo sit id. Minima nam magni dicta ut. Sit a laudantium consequatur natus modi molestias. Aliquid quos dolorem illo nesciunt et."}]},{"children":[{"text":""}]}]`,
		},
		models.Snippet{
			UserID:      5,
			DeckID:      5,
			Title:       "Super Bowl Predictions",
			Description: `[{"children":[{"text":"Sit ut fugit et quae. Velit rem esse molestias rerum. Animi ut animi. Doloremque excepturi molestias adipisci nesciunt. Molestias occaecati dolores sed repudiandae mollitia. Officiis maxime a dolores quas."}]},{"children":[{"text":""}]},{"children":[{"text":"Quisquam accusantium voluptates quia officia accusamus nihil voluptas perspiciatis repellendus. Dolor sit adipisci minima amet hic consequatur explicabo sit id. Minima nam magni dicta ut. Sit a laudantium consequatur natus modi molestias. Aliquid quos dolorem illo nesciunt et."}]},{"children":[{"text":""}]}]`,
		},
		models.Snippet{
			UserID:      5,
			DeckID:      7,
			Title:       "Final Project Pacman Competition",
			Description: `[{"children":[{"text":"Molestias occaecati dolores sed repudiandae mollitia. Officiis maxime a dolores quas."}]},{"children":[{"text":""}]},{"children":[{"text":"Quisquam accusantium voluptates quia officia accusamus nihil voluptas perspiciatis repellendus. Dolor sit adipisci minima amet hic consequatur explicabo sit id. Minima nam magni dicta ut. Sit a laudantium consequatur natus modi molestias. Aliquid quos dolorem illo nesciunt et."}]},{"children":[{"text":""}]}]`,
		},
		models.Snippet{
			UserID:      5,
			DeckID:      6,
			Title:       "Chapter 9 L1 lecture with sound (fixed)",
			Description: `[{"children":[{"text":"Velit rem esse molestias rerum. Animi ut animi. Doloremque excepturi molestias adipisci nesciunt. Molestias occaecati dolores sed repudiandae mollitia. Officiis maxime a dolores quas."}]},{"children":[{"text":""}]},{"children":[{"text":"Quisquam accusantium voluptates quia officia accusamus nihil voluptas perspiciatis repellendus. Dolor sit adipisci minima amet hic consequatur explicabo sit id. Minima nam magni dicta ut. Sit a laudantium consequatur natus modi molestias. Aliquid quos dolorem illo nesciunt et."}]},{"children":[{"text":""}]}]`,
		},
	}

	var todos = []models.Todo{
		models.Todo{
			UserID:   5,
			DeckID:   6,
			Title:    "Eating Lunch",
			Deadline: time.Now().Add(time.Hour * 72),
		},
		models.Todo{
			UserID:   5,
			DeckID:   6,
			Title:    "Having Fun",
			Deadline: time.Now(),
		},
		models.Todo{
			UserID:   5,
			DeckID:   7,
			Title:    "Practice Bunch",
			Deadline: time.Now().Add(time.Hour * 48),
		},
	}

	var bookmarks = []models.Bookmark{
		models.Bookmark{
			UserID:      5,
			URL:         "https://www.newyorker.com/science/elements/from-bats-to-human-lungs-the-evolution-of-a-coronavirus",
			Title:       "From Bats to Human Lungs, the Evolution of a Coronavirus",
			Description: "SARS-CoV-2, which honed its viral genome for thousands of years, behaves like a monstrous mutant hybrid of all the coronaviruses that came before it.",
			Thumbnail:   "https://media.newyorker.com/photos/5e7a5bd1ace85e0008863167/16:9/w_1280,c_limit/Kormann-Virusprofile-Respiratory.jpg",
			WordCount:   24215,
		},
		models.Bookmark{
			UserID:      5,
			URL:         "https://trix-editor.org/",
			Title:       "A rich text editor for everyday writing",
			Description: "Compose beautifully formatted text in your web application. Trix is an editor for writing messages, comments, articles, and lists—the simple documents most web apps are made of. It features a sophisticated document model, support for embedded attachments, and outputs terse and consistent HTML.",
			Thumbnail:   "",
			WordCount:   399,
		},
		models.Bookmark{
			UserID:      3,
			URL:         "https://medium.com/towards-artificial-intelligence/keras-callbacks-explained-in-three-minutes-846a43b44a16",
			Title:       "Keras Callbacks Explained In Three Minutes",
			Description: "A gentle introduction to callbacks in Keras. Learn about EarlyStopping, ModelCheckpoint, and other callback functions with code examples.",
			Thumbnail:   "https://miro.medium.com/max/1200/1*wwnExqe720PPHykHhs5Hqw.png",
			WordCount:   22712,
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

	for _, snippet := range snippets {
		if err := s.Snippet.Create(&snippet); err != nil {
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
