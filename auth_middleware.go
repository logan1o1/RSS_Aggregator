package main

import (
	"fmt"
	"net/http"

	"github.com/logan1o1/RSS_Aggregator/internal/auth"
	"github.com/logan1o1/RSS_Aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) authMiddleware(handler authHandler) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		apiKey, err := auth.GetAPIKey(req.Header)
		if err != nil {
			respondWithError(resp, 403, fmt.Sprintf("Error getting the api_key from the header: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUsersByAPIKey(req.Context(), apiKey)
		if err != nil {
			respondWithError(resp, 400, fmt.Sprintf("Error getting the user: %v", err))
			return
		}
		handler(resp, req, user)
	}
}
