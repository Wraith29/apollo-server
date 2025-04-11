package api

import (
	"errors"
	"net/http"
)

func RateRecommendation(w http.ResponseWriter, req *http.Request) {
	writeError(w, http.StatusInternalServerError, errors.New("Not Implemented"))
}
