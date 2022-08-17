package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "go-study/github-issues/github"
)

func main () {
    result := new(github.IssuesSearchResult) // contains current issues page
    var navStr string                        // contains links to previous, next, and last page

    navStr, result, err := github.SearchIssues(os.Args[1:], result)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(navStr)

    fmt.Printf("%d issues:\n", result.TotalCount)
    for _, item := range result.Items {
        fmt.Printf(
            "#%-5d %v %9.9s %.55s\n",
            item.Number, setTimeTranche(item.CreatedAt), item.User.Login, item.Title,
        )
    }
}

func setTimeTranche (t time.Time) string {
    now := time.Now()
    monthAgo := now.AddDate(0, -1, 0)
    yearAgo := now.AddDate(-1, 0, 0)

    if t.After(monthAgo) {
        return "Less than a month old"
    } else if t.After(yearAgo) {
        return "Less than a year old "
    } else {
        return "Older than a year    "
    }
}
