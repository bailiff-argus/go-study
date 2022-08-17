package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "bufio"

    "go-study/github-issues/github"
    "go-study/github-issues/parse"
)

func main () {
    result := new(github.IssuesSearchResult) // contains current issues page
    reader := bufio.NewReader(os.Stdin)

    // Mainloop
    for {
        navStr, result, err := github.SearchIssues(os.Args[1:], result)
        if err != nil {
            log.Fatal(err)
        }

        parse.ParseNavigation(navStr)

        fmt.Printf("%d issues:\n", result.TotalCount)
        for _, item := range result.Items {
            fmt.Printf(
                "#%-5d %v %9.9s %.55s\n",
                item.Number, setTimeTranche(item.CreatedAt), item.User.Login, item.Title,
            )
        }

        // Processing user input
        fmt.Printf("\nEnter command:\n:")
        input, err := reader.ReadString('\n')
        input = input[:len(input)-1]    // remove delimiter

        if err != nil {
            log.Fatal(err)
        }

        if (input == "Q") || (input == "q") {
            break
        }
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
