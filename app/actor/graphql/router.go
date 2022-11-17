package graphql

import (
	"net/http"

	"okp4/nemeton-leaderboard/graphql"
	"okp4/nemeton-leaderboard/graphql/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func makeRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/graphiql", playground.Handler("GraphiQL", "/graphql"))
	router.Handle(
		"/graphql",
		handler.NewDefaultServer(
			generated.NewExecutableSchema(
				generated.Config{Resolvers: &graphql.Resolver{}},
			),
		),
	)

	return router
}
