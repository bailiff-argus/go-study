package omdb

import (
	"fmt"
	"net/http"
	"net/url"
    "encoding/json"
)

func GetPosterLink(name, auth string) (string, error) {
    safeName := url.QueryEscape(name)
    reqLink := baseLink + fmt.Sprintf("?apikey=%s&t=%s", auth, safeName)

    client := http.Client{}

    req, err := http.NewRequest("GET", reqLink, nil)
    if err != nil { return "", err }

    resp, err := client.Do(req)
    if err != nil { return "", err }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK { return "", fmt.Errorf(resp.Status) }

    poster := poster{}
    if err := json.NewDecoder(resp.Body).Decode(&poster); err != nil {
        return "", err
    }

    return poster.ImgLink, nil
}
