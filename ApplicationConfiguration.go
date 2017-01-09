package main

import (
	"flag"
	"errors"
)

type ApplicationConfiguration struct {
	PodcastURI string
}

func buildApplicationConfiguration() (config ApplicationConfiguration, err error) {
	config = ApplicationConfiguration{}
	err = nil

	flag.StringVar(&config.PodcastURI, "podcast", "", "The URL of your podcast (i.E. https://feeds.example.com/example.m3u)")
	flag.Parse()
	if(config.PodcastURI == "") {		
		flag.Usage()
		err = errors.New("Please provide a valid podcast URI.")
	}
	
	return config, err
}
