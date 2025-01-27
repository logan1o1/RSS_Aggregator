package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
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
		UpdatedIt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(resp, 400, fmt.Sprintf("Err creating user in user_handler: %v", err))
		return
	}

	respondWithJSON(resp, 200, databaseUserToUser(user))
}
