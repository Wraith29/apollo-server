package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wraith29/apollo/internal/ctx"
)

var (
	invalidToken = errors.New("invalid token")
	expiredToken = errors.New("token has expired")
)

func parseAuthToken(header string) (*jwt.Token, error) {
	claims := jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(header, &claims, func(tkn *jwt.Token) (any, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidToken
		}

		//TODO: Replace this with an actual secret key
		return []byte(os.Getenv("APOLLO_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, invalidToken
	}

	expiry, err := claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	if expiry.Before(time.Now()) {
		return nil, expiredToken
	}

	return token, nil
}

func AuthenticationMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			writeError(w, http.StatusUnauthorized, errors.New("missing authorization header"))
			return
		}

		token, err := parseAuthToken(authHeader)
		if err != nil && (errors.Is(err, invalidToken) || errors.Is(err, expiredToken)) {
			fmt.Printf("%+v\n", err)
			writeError(w, http.StatusUnauthorized, err)
			return
		} else if err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}

		userId, err := token.Claims.GetSubject()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}

		userContext := context.WithValue(req.Context(), ctx.ContextKeyUserId, userId)

		next.ServeHTTP(w, req.WithContext(userContext))
	})
}
