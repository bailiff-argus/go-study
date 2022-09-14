package main

import (
    "os"
    "log"
	"net/http"

	"golang.org/x/net/html"
    "go-study/html-prettyprinter/print"
)

func main() {
    log.SetPrefix("outline2: ")

    resp, err := http.Get(os.Args[1])
    if err != nil { log.Fatal(err) }

    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Fatal(resp.Status)
    }

    doc, err := html.Parse(resp.Body)
    if err != nil { log.Fatal(err) }

    print.ForEachNode(doc, print.StartElement, print.EndElement)
}

