package api

import (
	"fmt"
	"github.com/wraith29/apollo/internal/ctx"
	"net/http"
)

func UpdateArtists(state *State, w http.ResponseWriter, req *http.Request) {
	userId := req.Context().Value(ctx.ContextKeyUserId).(string)

	fmt.Printf("%s -> %+v\n", userId, state)
}
