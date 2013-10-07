package importer

import (
	"github.com/julianduniec/news/rss"
	"github.com/julianduniec/news/store"
	"time"
	"fmt"
	"sync"
)

var (
	/*
		Contains all feeds
	*/
	feeds []rss.Feed
)

func Start(pollingFrequency time.Duration, rssFile string) {
	setupFeeds(rssFile)
	for ; ; {
		fmt.Println("import")
		startTime := time.Now()
		importFeeds()
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
	
	var wg sync.WaitGroup

	for _, uri := range feedUris {
		/*
			add feed async, pass waitgroup that allows
			addFeed to notify when done
		*/
		wg.Add(1)
		go addFeed(uri, &wg)
	}
	/*
		Await all the addFeed-goroutines to finish
	*/
	wg.Wait()
}


func addFeed(uri string, wg * sync.WaitGroup) {
	defer wg.Done()
	feed, err := rss.Fetch(uri)
	if err != nil {
		fmt.Println(err, uri)
		return
	}
	feeds = append(feeds, *feed)
}

func importFeeds() {
	var wg sync.WaitGroup
	for _, feed := range feeds {
		wg.Add(1)
		go importFeed(feed, &wg)
	}
	wg.Wait()
}

func importFeed(feed rss.Feed, wg *sync.WaitGroup) {
	defer wg.Done()
	feed.Update()
	for _, item := range feed.Items {
		importFeedItem(item)		
	}
	
}

func importFeedItem (item * rss.Item) {
	news := store.News{item.Link, item.Title, item.Content, "", item.Date}
	store.Add(news)
}
