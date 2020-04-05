package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/phamstack/godek/graph/generated"
	"github.com/phamstack/godek/graph/resolvers"
	"github.com/phamstack/godek/lib/auth"
	"github.com/phamstack/godek/lib/db"
	"github.com/phamstack/godek/lib/helpers"
	"github.com/phamstack/godek/models"
	"github.com/rs/cors"
)

func main() {
	port := "8088"
	connectionInfo := db.GetConnectionInfo()

	// connecting to postgres database
	// db, err := gorm.Open("postgres", connectionInfo)
	services, err := models.NewServices(
		models.WithGorm("postgres", connectionInfo),
		models.WithLogMode(true),
		models.WithUser(),
		models.WithDeck(),
		models.WithSnippet(),
		models.WithTodo(),
		models.WithBookmark(),
	)
	helpers.Must(err)
	services.DestructiveReset()
	defer services.Close()

	// seeds users
	db.SeedDatabase(services)

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:7877"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}).Handler)

	router.Use(auth.Middleware(services))
	router.Use(middleware.Logger)
	// initializing graphql server
	rootResolver := &resolvers.Resolver{
		Services: services,
	}
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: rootResolver}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	panic(http.ListenAndServe(":"+port, router))
}
