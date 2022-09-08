package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func BuildIndex(filename string) error {
    var searchRes []Comic

    for i := 1; i < 3000; i++ {
        if i == 404 { // very funny isn't it
            continue
        }
        comicDesc, status, err := fetchComic(i)
        if err != nil { return err }

        fmt.Println(comicDesc.Number, status)

        if status == http.StatusNotFound { break }

        searchRes = append(searchRes, comicDesc)
    }

    err := writeToJSON(searchRes, filename)
    return err
}

func writeToJSON(data []Comic, filename string) error {
    j, err := json.MarshalIndent(data, "", "  ")
    if err != nil { return err }

    err = ioutil.WriteFile(filename, j, 0644)
    return err
}

func fetchComic(number int) (Comic, int, error) {
    client    := http.Client{}
    comicDesc := Comic{}

    link := fmt.Sprintf(
        "%s%d%s",
        baseLinkLeft, number, baseLinkRight,
    )

    req, err := http.NewRequest("GET", link, nil)
    if err != nil { return Comic{}, 0, err }

    resp, err := client.Do(req)
    if err != nil { return Comic{}, 0, err }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return Comic{}, resp.StatusCode, nil
    }

    if err := json.NewDecoder(resp.Body).Decode(&comicDesc); err != nil {
        return Comic{}, resp.StatusCode, err
    }

    return comicDesc, resp.StatusCode, nil
}
