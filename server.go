package main

import (
	"github.com/phamstack/godek/graph"
	"github.com/phamstack/godek/postgres"
)

const defaultPort = "8088"

func main() {
	postgres.InitDB()

	graph.CreateGraphQLServer(defaultPort)
}
