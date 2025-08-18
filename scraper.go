package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Devansh-Awasthi/rssagg/internal/database"
	"github.com/google/uuid"
)

func parsePubDate(pubDate string) (time.Time, error) {
	layouts := []string{
		time.RFC1123Z, // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC1123,  // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC3339,  // "2006-01-02T15:04:05Z07:00"
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, pubDate)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unsupported date format: %s", pubDate)
}
func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("scrapping from %v goroutine every %s second", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Printf("err fetching feeds")
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}

}
func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.UpdateNextFeedToFetch(context.Background(), feed.ID)
	if err != nil {
		log.Printf("err updatinging feeds %s", err)
	}
	rssfeed, err := urlTofeed(feed.Url)
	if err != nil {
		log.Printf("err updatinging feeds %s", err)
	}

	for _, item := range rssfeed.Channel.Items {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		t, err := parsePubDate(item.PubDate)
		if err != nil {
			log.Printf("cant parse time")
		}
	  _,err =	db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: t,
			Url:         item.Link,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(),"duplicate key"){
				continue
			}
			log.Printf("err in creating post")
		}
	}
	log.Printf("feed %s collected ,%v post feed", feed.Name, len(rssfeed.Channel.Items))

}
