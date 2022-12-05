package graphql

import (
	"context"
	"net/http"
	"strings"

	"okp4/nemeton-leaderboard/graphql"
)

func NewBearerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		auth := request.Header.Get("Authorization")
		if len(auth) > 6 && strings.ToLower(auth[:7]) == "bearer " {
			bearer := auth[7:]
			ctx := context.WithValue(request.Context(), graphql.CTXBearerKey, bearer)

			request = request.WithContext(ctx)
		}

		next.ServeHTTP(writer, request)
	})
}
