package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wraith29/apollo/internal/ctx"
	"github.com/wraith29/apollo/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *authRequest) fromBody(body io.ReadCloser) error {
	if err := json.NewDecoder(body).Decode(&a); err != nil {
		return err
	}

	if a.Username == "" {
		return errors.New("missing required field: \"username\"")
	}

	if a.Password == "" {
		return errors.New("missing required field: \"password\"")
	}

	return nil
}

type authResponse struct {
	AuthToken string `json:"authToken"`
}

func Post_Register(w http.ResponseWriter, req *http.Request) {
	var ar authRequest
	if err := ar.fromBody(req.Body); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	usernameTaken := db.UsernameTaken(ar.Username)
	if usernameTaken {
		writeError(w, http.StatusConflict, fmt.Errorf("username \"%s\" is taken", ar.Username))
		return
	}

	hashedPassword, err := hashPassword(ar.Password)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	userId, err := db.SaveUser(ar.Username, hashedPassword)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	authToken, err := createAuthToken(userId)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if err := writeAuthResponse(w, authToken); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)

	return string(bytes), err
}

func Post_Login(w http.ResponseWriter, req *http.Request) {
	var ar authRequest
	if err := ar.fromBody(req.Body); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	user, err := db.GetUserByUsername(ar.Username)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ar.Password)); err != nil {
		writeError(w, http.StatusUnauthorized, errors.New("invalid username or password"))
		return
	}

	authToken, err := createAuthToken(user.Id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if err := writeAuthResponse(w, authToken); err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func Get_Refresh(w http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value(ctx.ContextKeyUserId).(string)

	newToken, err := createAuthToken(userId)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	if err := writeAuthResponse(w, newToken); err != nil {
		fmt.Printf("%+v\n", err)
	}
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

func writeAuthResponse(w http.ResponseWriter, authToken string) error {
	response := authResponse{AuthToken: authToken}
	body, err := json.Marshal(&response)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return nil
	}

	w.Header().Add("Content-Type", "application/json")
	if _, err := w.Write(body); err != nil {
		return err
	}

	return nil
}
