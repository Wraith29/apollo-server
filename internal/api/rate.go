package api

import (
	"errors"
	"net/http"
)

func Put_Rating(w http.ResponseWriter, req *http.Request) {
	writeError(w, http.StatusInternalServerError, errors.New("Not Implemented"))
}
