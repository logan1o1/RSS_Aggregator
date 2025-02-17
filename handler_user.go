package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	// "github.com/logan1o1/RSS_Aggregator/internal/auth"
	"github.com/logan1o1/RSS_Aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(resp http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(req.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(resp, 400, fmt.Sprintf("Err parsing json in user_handler: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(req.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(resp, 400, fmt.Sprintf("Err creating user in user_handler: %v", err))
		return
	}

	respondWithJSON(resp, 201, databaseUserToUser(user))
}

// user database.User

func (apiCfg *apiConfig) handlerGetUser(resp http.ResponseWriter, req *http.Request) {
	type parameters struct {
		ApiKey string `json:"api_key"`
	}
	decoder := json.NewDecoder(req.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(resp, 400, fmt.Sprintf("Err parsing json in get_user_handler: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUsersByAPIKey(req.Context(), params.ApiKey)
	if err != nil {
		respondWithError(resp, 400, fmt.Sprintf("Err getting user in get_user_handler: %v", err))
		return
	}

	respondWithJSON(resp, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUsers(resp http.ResponseWriter, req *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostForUser(req.Context(), database.GetPostForUserParams{
		Userid: user.ID,
		Limit:  5,
	})
	if err != nil {
		respondWithError(resp, 400, fmt.Sprintf("Err getting posts for users user_handler: %v", err))
		return
	}
	respondWithJSON(resp, 200, databasePostsToPosts(posts))
}
