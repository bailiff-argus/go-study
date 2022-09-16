package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"go-study/webcrawler/links"
	"go-study/webcrawler/search"
)

func main() {
    search.BreadthFirst(crawl, os.Args[1:])
    // search.DepthFirst(crawl, os.Args[1:])
}

var domain string

func crawl(a string) []string {
    log.Printf("webcrawler: crawling %s\n", a)
    // https:// + link.Host + link.Path

    link, err := url.Parse(a)
    if err != nil { log.Printf("webcrawler: %s", err) }

    if domain == "" {
        domain = link.Host
    }

    if link.Host == domain {
        dir   := "./temp/" + link.Host + link.Path
        fName := "temp.html"
        if !strings.HasSuffix(dir, "/") { fName = "/" + fName }

        err := createLocalCopy(a, dir, fName)
        if err != nil { log.Printf("webcrawler: %s", err) }
    } 

    list, err := links.Extract(a)
    if err != nil { log.Println(err) }

    return list
}

func createLocalCopy(a, folderName, fileName string) error {
    pgContent, err := getPage(a)
    if err != nil { return err }

    err = writeContentsToFile(pgContent, folderName, fileName)
    if err != nil { return err }

    return nil
}

func writeContentsToFile(text, dir, fileName string) error {
    err := os.MkdirAll(dir, 0766)
    if err != nil { return err }

    file, err := os.Create(dir + fileName)
    if err != nil { return err }

    defer file.Close()

    _, err  = file.WriteString(text)
    if err != nil { return err }

    log.Printf("webcrawler: %s writing success", file.Name())

    return nil
}

func getPage(a string) (string, error) {
    resp, err := http.Get(a)
    if err != nil { return "", err }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("cannot fetch %s url: %s", a, resp.Status)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil { return "", err }

    return string(body), nil
}
