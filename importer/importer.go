package importer

import (
	"github.com/SlyMarbo/rss"
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
	/*
		Used to lock the feeds-collection
	*/
	feedLock sync.Mutex
)

func Start(pollingFrequency time.Duration, rssFile string) {
	setupFeeds(rssFile)
	for ; ; {
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
	feed, err := rss.Fetch(uri)
	if err != nil {
		fmt.Println(err, uri)
		return
	}
	
	feedLock.Lock()
	feeds = append(feeds, *feed)
	feedLock.Unlock()
	
	wg.Done()
}

func importFeeds() {
	var wg sync.WaitGroup
	feedLock.Lock()
	for _, feed := range feeds {
		wg.Add(1)
		go importFeed(feed, &wg)
	}
	feedLock.Unlock()
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
