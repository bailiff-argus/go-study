package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"go-study/xkcd/ctrl"
	"go-study/xkcd/xkcd"
    "go-study/github-issues/parse"
)

func main() {
    dbName := "xkcddb.json"
    forceRebuld, noUpdate, searchTerm := ctrl.ReadFlags()

    if (dbNotExistOrEmpty(dbName) || forceRebuld) && !noUpdate {
        err := xkcd.BuildIndex(dbName)
        if err != nil {
            log.Printf("db rebuild error: %s", err)
        }
    }

    searchRes, err := SearchInDB(dbName, searchTerm)
    if err != nil { log.Fatal(err) }

    searchResStr := printSearchResults(searchRes)
    parse.ShowInPager(searchResStr)
}

func dbNotExistOrEmpty(filename string) bool {
    desc, err := os.Stat(filename)
    if errors.Is(err, os.ErrNotExist) {
        return true
    }

    if desc.Size() <= 1048576 {
        return true
    }

    return false
}

func printSearchResults(db []xkcd.Comic) string {
    var text string
    for _, comic := range db {
        comicStr := fmt.Sprintf(
            "###NEXT COMIC###\nTitle: %s\nDescr: %s\nURL:   %s\n\n\n\n",
            comic.Title, comic.Transcript, comic.URL,
        )

        text += comicStr
    }

    return text
}
