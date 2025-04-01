package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wraith29/apollo/internal/data"
)

type initRequest struct {
	Username string `json:"username"`
}

func Init(w http.ResponseWriter, req *http.Request) {
	var body initRequest
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	hash := sha256.New()
	if _, err := hash.Write([]byte(body.Username)); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	userId := hex.EncodeToString(hash.Sum(nil))[:8]

	if err := data.SaveUser(userId, body.Username); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(userId)); err != nil {
		fmt.Println(err)
	}
}
