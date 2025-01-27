package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/logan1o1/RSS_Aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedIt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedIt: dbUser.UpdatedIt,
		Name:      dbUser.Name,
	}
}
