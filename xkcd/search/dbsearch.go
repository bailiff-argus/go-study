package main

import (
    "strings"
    "io/ioutil"
    "encoding/json"

    "go-study/xkcd/xkcd"
)

func SearchInDB(dbName, term string) ([]xkcd.Comic, error) {
    dbContent, err := fetchDB(dbName)
    if err != nil { return []xkcd.Comic{}, err }

    searchRes := findInDB(dbContent, term)

    return searchRes, nil
}

func fetchDB(dbName string) ([]xkcd.Comic, error) {
    var db []xkcd.Comic

    dbFile, err := ioutil.ReadFile(dbName)
    if err != nil { return []xkcd.Comic{}, err }

    if err := json.Unmarshal(dbFile, &db); err != nil { return []xkcd.Comic{}, err }
    return db, nil
}

func findInDB(db []xkcd.Comic, term string) []xkcd.Comic {
    var result []xkcd.Comic

    terms := strings.Split(strings.ToLower(term), " ")
    for _, comic := range db {
        score := calcScore(comic, terms)
        if score >= 0.66 {
            result = append(result, comic)
        }
    }

    return result
}

func calcScore(comic xkcd.Comic, terms []string) float32 {
    fullText  := strings.ToLower(comic.Title) + " " + strings.ToLower(comic.Transcript)
    fullWords := strings.Split(fullText, " ")
    maxScore  := float32(len(terms))
    currScore := float32(0)

    for _, term := range terms {
        if contains(fullWords, term) {
            currScore += 1.0
            continue
        }
    }

    return currScore / maxScore
}

func contains(set []string, str string) bool {
    for _, item := range set {
        if str == item {
            return true
        }
    }

    return false
}
