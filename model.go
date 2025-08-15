package main

import (
	"time"

	"github.com/Devansh-Awasthi/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	APIKey    string    `json:"api_key"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string	`json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	Url       string	`json:"url"`
	UserID    uuid.UUID	`json:"user_id"`
}
func databaseFeedtoFeed(dbfeed database.Feed) Feed {
	return Feed{
		ID:        dbfeed.ID,
		Name:      dbfeed.Name,
		CreatedAt: dbfeed.CreatedAt,
		UpdatedAt: dbfeed.UpdatedAt,
		Url: dbfeed.Url,
		UserID: dbfeed.UserID,
	}
}
func databaseFeedstoFeeds(dbfeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _,dbfeed := range dbfeeds {
		feeds = append(feeds, databaseFeedtoFeed(dbfeed))
	} 
	return feeds
}
type FeedFollows struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	UserID    uuid.UUID	`json:"user_id"`
	FeedID    uuid.UUID	`json:"feed_id"`
}
func databaseFeedFolltoFeedFoll(dbfeed database.FeedsFollowing) FeedFollows {
	return FeedFollows{
		ID:        dbfeed.ID,
		CreatedAt: dbfeed.CreatedAt,
		UpdatedAt: dbfeed.UpdatedAt,
		UserID: dbfeed.UserID,
		FeedID : dbfeed.FeedID,
	}
}
func databaseGetFeedstoGetFeeds(dbfeeds []database.FeedsFollowing) []FeedFollows {
	feeds := []FeedFollows{}
	for _,dbfeed := range dbfeeds {
		feeds = append(feeds, databaseFeedFolltoFeedFoll(dbfeed))
	} 
	return feeds
}