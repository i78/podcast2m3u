package main

import (
	"fmt"

	"github.com/nandosousafr/podfeed"
)

func M3u(podcast podfeed.Podcast) {
	fmt.Println("#EXTM3U")
	for _, episode := range podcast.Items {
		fmt.Printf("#EXTINF:%d,%s\n%s\n",
			0,
			episode.Title,
			episode.Enclosure.Url)
	}

}
