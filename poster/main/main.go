package main

import (
	"log"

	"go-study/poster/ctrl"
	"go-study/poster/omdb"
	"go-study/poster/poster"
)


func main() {
    movieName, authToken, outFileName := ctrl.GetFlags()

    if movieName == "" {
        log.Fatal("poster: no movie name specified")
    }

    if authToken == "" {
        log.Fatal("poster: no auth token specified")
    }

    posterLink, err := omdb.GetPosterLink(movieName, authToken)
    if err != nil { log.Fatal(err) }

    if err := poster.SaveImageToFile(posterLink, outFileName); err != nil {
        log.Fatal(err)
    }
}
