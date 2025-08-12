package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Devansh-Awasthi/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreatefeed(w http.ResponseWriter, r *http.Request,user database.User) {
	type parameter struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	param := parameter{}
	err := decoder.Decode(&param)
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintf("Error on Parsing:%s", err))
		return
	}
		feed,err := apiCfg.db.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      param.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url: param.Url,
		UserID: user.ID,
	})
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintf("Error on Creating User:%s", err))
		return
	}
	responseWithJSON(w, 200, databaseFeedtoFeed(feed))
}
func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	
		feeds,err := apiCfg.db.GetFeeds(r.Context())
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintf("Error on Creating User:%s", err))
		return
	}
	responseWithJSON(w, 200, databaseFeedstoFeeds(feeds))
}
