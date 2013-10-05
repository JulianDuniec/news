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
	fmt.Println(pollingFrequency)
	setupFeeds(rssFile)
	for _ = range time.Tick(pollingFrequency) {
		if finished == true {
			doImport()
		}
		
	}
}

func setupFeeds(rssFile string) {

	feedUris, err := fetchRssList(rssFile)
	if err != nil {
		fmt.Println("Could not open file", err)
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
	finished = false
	fmt.Println("importer:doImport()")
	var wg sync.WaitGroup
	for _, feed := range feeds {
		feed.Update()
		for _, item := range feed.Items {
			wg.Add(1)
			go importFeedItem(item, &wg)		
		}
	}
	fmt.Println("importer:doImport():waiting")
	wg.Wait()
	fmt.Println("importer:doImport():done")
	finished = true
}

func importFeedItem (item * rss.Item, wg *sync.WaitGroup) {
	news := store.News{item.Link, item.Title, item.Content, "", item.Date}
	store.Add(news)
	fmt.Println(item.Link)
	wg.Done()

}
