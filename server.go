package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/phamstack/godek/graph/generated"
	"github.com/phamstack/godek/graph/resolvers"
	"github.com/phamstack/godek/helpers"
	"github.com/phamstack/godek/models"
)

func main() {
	port := "8088"
	connectionInfo := helpers.GetConnectionInfo()

	// connecting to postgres database
	services, err := models.NewServices(
		models.WithGorm("postgres", connectionInfo),
		models.WithLogMode(true),
		models.WithUser(),
	)
	helpers.Must(err)
	services.DestructiveReset()
	defer services.Close()

	// router := chi.NewRouter()
	// router.Use(auth.Middleware(services))
	// initializing graphql server
	rootResolver := &resolvers.Resolver{
		Services: services,
	}
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: rootResolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
