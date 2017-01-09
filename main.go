package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nandosousafr/podfeed"
	"m9d.de/podcast2m3u/exportfilters"
)

func main() {
	conf, confError := buildApplicationConfiguration()
	if confError != nil {
		log.Fatal(confError)
		os.Exit(1)
	} else {
		podcast, err := podfeed.Fetch(conf.PodcastURI)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(podcast.Title)
		exportfilters.M3u(podcast)
	} 
}
