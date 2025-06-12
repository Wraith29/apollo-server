package api

import (
	"fmt"
	"net/http"
)

func writeError(w http.ResponseWriter, status int, err error) {
	bytes := fmt.Appendf(nil, `{"err": "%s"}`, err.Error())

	w.WriteHeader(status)

	if _, err := w.Write(bytes); err != nil {
		fmt.Printf("error responding: %+v\n", err)
	}
}
