package api

import (
	"errors"
	"net/http"
)

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, req *http.Request) {
	writeError(w, http.StatusInternalServerError, errors.New("Not Implemented"))
}

func Login(w http.ResponseWriter, req *http.Request) {
	writeError(w, http.StatusInternalServerError, errors.New("Not Implemented"))
}

func Refresh(w http.ResponseWriter, req *http.Request) {
	writeError(w, http.StatusInternalServerError, errors.New("Not Implemented"))
}
