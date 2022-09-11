package main

import (
	"fmt"
	"log"
	"os"

	"net/http"
	"net/url"

	"image/jpeg"

	"encoding/json"

	"go-study/poster/ctrl"
)

const baseLink string = "https://www.omdbapi.com/"

type Poster struct{
    ImgLink     string      `json:"Poster"`
}   

func main() {
    movieName, authToken, output := ctrl.GetFlags()

    outFile, err := os.Create(output)
    if err != nil { log.Fatal(err) }
    defer outFile.Close()

    if movieName == "" {
        log.Fatal("poster: no movie name specified")
    }

    if authToken == "" {
        log.Fatal("poster: no auth token specified")
    }

    safeMovieName := url.QueryEscape(movieName)
    pars := fmt.Sprintf("?apikey=%s&t=%s", authToken, safeMovieName)
    request := baseLink + pars
    log.Println(request)

    client := http.Client{}
    req, err := http.NewRequest("GET", request, nil)
    if err != nil { log.Fatal(err) }

    resp, err := client.Do(req)
    if err != nil { log.Fatal(err) }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK { log.Fatal(resp.Status) }

    poster := Poster{}
    if err := json.NewDecoder(resp.Body).Decode(&poster); err != nil {
        log.Fatal(err)
    }

    req, err = http.NewRequest("GET", poster.ImgLink, nil)
    if err != nil { log.Fatal(err) }

    resp, err = client.Do(req)
    if err != nil { log.Fatal(err) }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK { log.Fatal(resp.Status) }

    image, err := jpeg.Decode(resp.Body)
    if err != nil { log.Fatal(err) }

    if err := jpeg.Encode(outFile, image, nil); err != nil {
        log.Fatal(err)
    }
}
