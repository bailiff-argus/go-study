package main

import (
    "os"
	"fmt"
	"log"

	"go-study/webcrawler/links"
	"go-study/webcrawler/search"
)

func main() {
    search.BreadthFirst(crawl, os.Args[1:])
}

func crawl(url string) []string {
    fmt.Println(url)
    list, err := links.Extract(url)
    if err != nil { log.Println(err) }

    return list
}
