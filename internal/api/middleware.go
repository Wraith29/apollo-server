package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/wraith29/apollo/internal/ctx"
)

func UserIdMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		userIdHeader := req.Header.Get("User-Id")
		if userIdHeader == "" {
			writeError(w, http.StatusUnauthorized, errors.New("Missing required header: \"User-Id\""))
			return
		}

		userIdContext := context.WithValue(req.Context(), ctx.ContextKeyUserId, userIdHeader)

		next.ServeHTTP(w, req.WithContext(userIdContext))
	})
}
