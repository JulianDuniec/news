package importer

import (
	"github.com/SlyMarbo/rss"
	"github.com/julianduniec/news/store"
	"time"
)

var (
	feeds []rss.Feed
	finished					= true
	pollingFrequencySeconds 	= 5 * time.Second
	feedUris 					= []string {
		"http://www.aftonbladet.se/nyheter/rss.xml",
		"http://www.aftonbladet.se/sportbladet/rss.xml",
		"http://www.aftonbladet.se/nojesbladet/rss.xml",
		"http://www.aftonbladet.se/kultur/rss.xml"}
)

func Start() {
	setupFeeds()
	for _ = range time.Tick(pollingFrequencySeconds) {
		doImport()
	}
}

func setupFeeds() {
	for _, uri := range feedUris {
		feed, _ := rss.Fetch(uri)
		feeds = append(feeds, *feed)
	}
}

func doImport() {
	for _, feed := range feeds {
		feed.Update()
		for _, item := range feed.Items {
			importFeedItem(item)
		}
	}
}

func importFeedItem (item * rss.Item) {
	news := store.News{item.Link, item.Title, item.Content, ""}
	store.Add(news)
}
