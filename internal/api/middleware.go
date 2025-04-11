package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
		return []byte("SecretKey"), nil
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
		authHeader := req.Header.Get("Authentication")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := parseAuthToken(authHeader)
		if err != nil && (errors.Is(err, invalidToken) || errors.Is(err, expiredToken)) {
			writeError(w, http.StatusUnauthorized, err)
			return
		} else if err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}

	})
}
