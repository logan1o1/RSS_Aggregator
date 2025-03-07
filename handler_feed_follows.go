package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/logan1o1/RSS_Aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(resp http.ResponseWriter, req *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feedid"`
	}
	decoder := json.NewDecoder(req.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(resp, 400, fmt.Sprintf("Err parsing json in feed_follow_handler: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(req.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Userid:    user.ID,
		Feedid:    params.FeedID,
	})
	if err != nil {
		respondWithError(resp, 400, fmt.Sprintf("Err creating feedFollow in feed_follow_handler: %v", err))
		return
	}

	respondWithJSON(resp, 201, databaseFFollowToFFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(resp http.ResponseWriter, req *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(req.Context(), user.ID)
	if err != nil {
		respondWithError(resp, 400, fmt.Sprintf("Err getting feedFollows in feed_follow_handler: %v", err))
		return
	}

	respondWithJSON(resp, 200, databaseFFollowsToFFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(resp http.ResponseWriter, req *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(req, "feedFollowID")
	feedFollowId, err := uuid.Parse(feedFollowIdStr)
	if err != nil {
		respondWithError(resp, 400, fmt.Sprintf("Err parsing feedFollowStr in feed_follow_handler: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(req.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		Userid: user.ID,
	})
	if err != nil {
		respondWithError(resp, 400, fmt.Sprintf("Err deleting feedFollow in feed_follow_handler: %v", err))
		return
	}
	respondWithJSON(resp, 200, struct{}{})
}
