package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func makeRouter(graphqlServer *handler.Server) *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/graphiql", playground.Handler("GraphiQL", "/graphql"))
	router.Handle("/graphql", NewBearerMiddleware(graphqlServer))

	return router
}
