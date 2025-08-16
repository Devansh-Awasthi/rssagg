package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Devansh-Awasthi/rssagg/internal/database"
)

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
		log.Println("feed collected:", item.Title, "on Feed", feed.Name)
	}
	log.Printf("feed %s collected ,%v post feed", feed.Name, len(rssfeed.Channel.Items))

}
