package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Devansh-Awasthi/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreatefeedFollowing(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	param := parameter{}
	err := decoder.Decode(&param)
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintf("Error on Parsing:%s", err))
		return
	}
	feed, err := apiCfg.db.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    param.FeedID,
	})
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintf("Error on Following feeds:%s", err))
		return
	}
	responseWithJSON(w, 200, databaseFeedFolltoFeedFoll(feed))
}
func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.db.GetFeedFollow(r.Context(), user.ID)
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintf("Error on Getting Feed :%s", err))
		return
	}
	responseWithJSON(w, 200, databaseGetFeedstoGetFeeds(feedFollows))
}
func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDstr := chi.URLParam(r, "feedFollowingID")
	feedFollowID, err := uuid.Parse(feedFollowIDstr)
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintf("Error on Following feeds:%s", err))
		return
	}
	err = apiCfg.db.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		responseWithErr(w, 400, fmt.Sprintf("Error on deleting feeds:%s", err))
		return
	}
	responseWithJSON(w, 200, struct{}{})
}
