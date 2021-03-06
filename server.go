package main

import (
	"log"
	"net/http"

	"github.com/mskarbe/go-gql-api-server/db/postgres"
	"github.com/mskarbe/go-gql-api-server/pkg/config"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	graph "github.com/mskarbe/go-gql-api-server/graph"
	"github.com/mskarbe/go-gql-api-server/graph/generated"
)

func main() {

	env := config.GetConfig()

	// database connect & migrate
	var db postgres.Postgres
	db.Connect(env.Db)
	db.Migrate()
	defer db.Close()

	resolver := &graph.Resolver{DbSchema: db.Database}
	resolver.Initialize()
	log.Println("Resolver initialized")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{ Resolvers: resolver }))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", env.Gql.Port)
	log.Fatal(http.ListenAndServe(":"+env.Gql.Port, nil))
}
