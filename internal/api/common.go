package api

import (
	"fmt"
	"net/http"
)

func writeError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	if _, err := w.Write([]byte(err.Error())); err != nil {
		fmt.Println(err.Error())
	}
}
