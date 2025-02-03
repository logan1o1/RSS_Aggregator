package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/logan1o1/RSS_Aggregator/internal/database"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetwnRequest time.Duration,
) {
	log.Printf("scraping on %v goroutines every %s duration", concurrency, timeBetwnRequest)
	ticker := time.NewTicker(timeBetwnRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("error fetching feeds", err)
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

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("error marking feed as fetched", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("error urlToFeed", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("found post", item.Title, "on feed", feed.Name)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
