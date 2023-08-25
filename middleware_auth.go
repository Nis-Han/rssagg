package main

import (
	"fmt"
	"net/http"

	"github.com/nerd500/rssagg/internal/auth"
	"github.com/nerd500/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (ApiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("auth Error: %v", err))
			return
		}

		user, err := ApiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, "invalid API Key")
			return
		}
		handler(w, r, user)
	}
}
