package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wraith29/apollo/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Post_Register(w http.ResponseWriter, req *http.Request) {
	var body authRequest

	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	usernameTaken := db.UsernameTaken(body.Username)
	if usernameTaken {
		writeError(w, http.StatusConflict, fmt.Errorf("username \"%s\" is taken", body.Username))
		return
	}

	hashedPassword, err := hashPassword(body.Password)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	userId, err := db.SaveUser(body.Username, hashedPassword)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	authToken, err := createAuthToken(userId)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	req.Header.Add("Authentication", fmt.Sprintf("Bearer %s", authToken))
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	return string(bytes), err
}

func Post_Login(w http.ResponseWriter, req *http.Request) {
	writeError(w, http.StatusInternalServerError, errors.New("Not Implemented"))
}

func Get_Refresh(w http.ResponseWriter, req *http.Request) {
	writeError(w, http.StatusInternalServerError, errors.New("Not Implemented"))
}

func createAuthToken(userId string) (string, error) {
	now := time.Now()

	// TODO: Make this configurable
	expiry := now.Add(time.Hour * 24 * 7)

	claims := jwt.RegisteredClaims{
		Issuer:    "apollo-server",
		Subject:   userId,
		Audience:  jwt.ClaimStrings{"http://localhost:5000"},
		ExpiresAt: jwt.NewNumericDate(expiry),
		IssuedAt:  jwt.NewNumericDate(now),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("APOLLO_SECRET_KEY")

	return token.SignedString([]byte(secretKey))
}
