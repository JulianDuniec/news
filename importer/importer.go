package importer

import (
	"github.com/SlyMarbo/rss"
	"github.com/julianduniec/news/store"
	"time"
	"fmt"
	"sync"
)

var (
	feeds []rss.Feed
	finished					= true
)

func Start(pollingFrequency time.Duration, rssFile string) {
	setupFeeds(rssFile)
	
	for ; ; {
		startTime := time.Now()
		doImport()
		/*
			Calculate the remainder of the pollingFrequency - the execution-time
			and re-execute function to keep a pace as close to polling-frequency as possible

			This allows doImport() to take longer than the pollingFrequency
		*/
		time.Sleep(pollingFrequency - time.Since(startTime))
	}
	
}

func setupFeeds(rssFile string) {
	feedUris, err := fetchRssList(rssFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %s", rssFile))
	}
	for _, uri := range feedUris {
		go addFeed(uri)
	}
}

func addFeed(uri string) {
	feed, err := rss.Fetch(uri)
	if err != nil {
		fmt.Println(err, uri)
		return
	}
	feeds = append(feeds, *feed)
}

func doImport() {
	var wg sync.WaitGroup

	for _, feed := range feeds {
		wg.Add(1)
		go importFeed(feed, &wg)
	}
	wg.Wait()
}

func importFeed(feed rss.Feed, wg *sync.WaitGroup) {
	feed.Update()
	for _, item := range feed.Items {
		importFeedItem(item)		
	}
	wg.Done()
}

func importFeedItem (item * rss.Item) {
	news := store.News{item.Link, item.Title, item.Content, "", item.Date}
	store.Add(news)

}
