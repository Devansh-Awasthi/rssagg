package main

import (
	"time"
	"github.com/Devansh-Awasthi/rssagg/internal/database"
	"github.com/google/uuid"
)
type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string  `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	APIKey string `json:"api_key"`
}
func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID  :      dbUser.ID,
	Name : dbUser.Name ,
	CreatedAt:dbUser.CreatedAt ,
	UpdatedAt : dbUser.UpdatedAt,
	APIKey : dbUser.ApiKey,
	}
}
