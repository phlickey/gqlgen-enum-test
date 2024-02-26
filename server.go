package main

import (
	"gqlgen-test/graph"
	"gqlgen-test/graph/model"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	todo1 := model.Todo{
		ID: "T1", Text: "test", User: &model.User{ID: "U1", Name: "user U1"}, Priority: model.PriorityMedium,
	}
	todo2 := model.Todo{
		ID: "T1", Text: "test", User: &model.User{ID: "U1", Name: "user U1"}, Priority: "🔥🔥🔥 FUEGO",
	}

	todos := []*model.Todo{
		&todo1, &todo2,
	}
	resolva := graph.NewResolver(todos)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolva}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
