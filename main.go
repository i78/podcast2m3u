package main

import (
	"fmt"
	"log"
	"os"
	"github.com/nandosousafr/podfeed"
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
		M3u(podcast)
	} 
}
