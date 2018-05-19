package youtubeApi

import (
	"net/http"
	"log"
	"flag"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
	query      = flag.String("query", "Google", "Search term")
	maxResults = flag.Int64("max-results", 3, "Max YouTube results")
	)
	
const developerKey = "ENTER YOUR API KEY HERE"

func YoutubeSearch(searchTerm *string) map[string]string {

	//log.Fatalf("searchTerm: %v", *searchTerm) 
	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	// Make the API call to YouTube.
	call := service.Search.List("id,snippet").
		Q(*searchTerm).
		MaxResults(*maxResults)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error retrieving response", err)
	}

	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
				videos[item.Id.VideoId] = item.Snippet.Title
		}
}

	return videos

	//printIDs("Videos", videos)
}
