package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/logan1o1/RSS_Aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(resp http.ResponseWriter, req *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(req.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(resp, 400, fmt.Sprintf("Err parsing json in feed_handler: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(req.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		Userid:    user.ID,
	})
	if err != nil {
		respondWithError(resp, 400, fmt.Sprintf("Err creating feed in feed_handler: %v", err))
		return
	}

	respondWithJSON(resp, 201, databaseFeedtoFeed(feed))
}

func (apicCfg *apiConfig) handlerGetFeeds(resp http.ResponseWriter, req *http.Request) {
	feeds, err := apicCfg.DB.GetFeeds(req.Context())
	if err != nil {
		respondWithError(resp, 400, fmt.Sprintf("Err getting feed in feed_handler: %v", err))
		return
	}
	respondWithJSON(resp, 200, databaseFeedstoFeeds(feeds))
}
