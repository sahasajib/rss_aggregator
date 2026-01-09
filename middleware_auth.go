package main

import (
	"fmt"
	"net/http"

	"github.com/sahasajib/rssagg/internal/auth"
	"github.com/sahasajib/rssagg/internal/database"
)

type autheadHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCnf *apiConfig) middlewareAuth(handler autheadHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apiCnf.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
