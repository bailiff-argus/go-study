package parse

import (
	"fmt"
	"time"
    "bufio"
    "strings"

	"go-study/github-issues/github"
)

func DisplayResult (result *github.IssuesSearchResult) {
    fmt.Printf("%d issues:\n", result.TotalCount)
    for _, item := range result.Items {
        fmt.Printf(
            "#%-5d %v %9.9s %.55s\n",
            item.Number, setTimeTranche(item.CreatedAt),
            item.User.Login, item.Title,
        )
    }

}

func ReceiveInput (reader *bufio.Reader) (string, error) {
    input, err := reader.ReadString('\n')
    if err != nil {
        return "", err
    }

    input = strings.ToLower(
        input[:len(input)-1],
    )

    return input, nil
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

