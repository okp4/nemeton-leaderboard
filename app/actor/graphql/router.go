package graphql

import (
	"net/http"

	"okp4/nemeton-leaderboard/graphql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func makeRouter(graphqlServer *handler.Server) *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/graphiql", playground.Handler("GraphiQL", "/graphql"))
	router.Handle("/graphql", NewBearerMiddleware(graphql.BearerCTXKey, graphqlServer))

	return router
}
