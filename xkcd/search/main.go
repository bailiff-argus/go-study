package main

import (
    "os"
    "errors"
    "log"
    "fmt"

    "go-study/xkcd/xkcd"
    "go-study/xkcd/ctrl"
)


func main() {
    dbName := "xkcddb.json"

    forceRebuld, searchTerm := ctrl.ReadFlags()
    fmt.Println(searchTerm)

    if dbNotExistOrEmpty(dbName) || forceRebuld {
        err := xkcd.BuildIndex(dbName)
        if err != nil {
            log.Printf("db rebuild error: %s", err)
        }
    }

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
