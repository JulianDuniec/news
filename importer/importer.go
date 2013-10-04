package importer

import (
	"github.com/SlyMarbo/rss"
	"github.com/julianduniec/news/store"
	"time"
	"fmt"
)

var (
	feeds []rss.Feed
	finished					= true
	feedUris 					= []string {
		"http://www.aftonbladet.se/nyheter/rss.xml",
		"http://www.aftonbladet.se/sportbladet/rss.xml",
		"http://www.aftonbladet.se/nojesbladet/rss.xml",
		"http://www.aftonbladet.se/kultur/rss.xml"}
)

func Start(pollingFrequency time.Duration) {
	fmt.Println(pollingFrequency)
	setupFeeds()
	for _ = range time.Tick(pollingFrequency) {
		doImport()
	}
}

func setupFeeds() {
	fmt.Println("importer:setupFeeds()")
	for _, uri := range feedUris {
		feed, _ := rss.Fetch(uri)
		feeds = append(feeds, *feed)
	}
}

func doImport() {
	fmt.Println("importer:doImport()")
	for _, feed := range feeds {
		feed.Update()
		for _, item := range feed.Items {
			importFeedItem(item)
		}
	}
}

func importFeedItem (item * rss.Item) {
	news := store.News{item.Link, item.Title, item.Content, "", item.Date}
	store.Add(news)
}
