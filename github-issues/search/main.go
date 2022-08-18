package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	// "os/exec"
	"bufio"
	"time"

	"go-study/github-issues/github"
	"go-study/github-issues/parse"
)

func main () {
    result := new(github.IssuesSearchResult) // contains current issues page
    var navStr string

    navStr, result, err := github.SearchIssues(os.Args[1:], result)
    if err != nil {
        log.Fatal(err)
    }

    reader := bufio.NewReader(os.Stdin)

    // Mainloop
    for {
        // cmd := exec.Command("clear")
        // cmd.Stdout = os.Stdout
        // cmd.Run()

        navigator := parse.ParseNavigation(navStr)

        fmt.Printf("%d issues:\n", result.TotalCount)
        for _, item := range result.Items {
            fmt.Printf(
                "#%-5d %v %9.9s %.55s\n",
                item.Number, setTimeTranche(item.CreatedAt), item.User.Login, item.Title,
            )
        }

        // Processing user input
        fmt.Printf("\nCOMMAND | :")
        input, err := reader.ReadString('\n')
        input = input[:len(input)-1]    // remove delimiter
        input = strings.ToLower(input)

        if err != nil {
            log.Fatal(err)
        }

        if input == "q" { // [q]uit
            break
        } else if strings.HasPrefix(input, "o") { // [o]pen

        } else { // if not those actions, then navigate
            var q string
            var ok bool

            if input == "n" { // [n]ext
                q, ok = navigator["next"]
                if !ok {
                    continue
                }
            } else if input == "p" { // [p]revious
                q, ok = navigator["prev"]
                if !ok {
                    continue
                }
            } else if input == "f" { // [f]irst
                q, ok = navigator["first"]
                if !ok {
                    continue
                }
            } else if input == "l" { // [l]ast
                q, ok = navigator["last"]
                if !ok {
                    continue
                }
            }

            navStr, result, err = github.SendRequest(q, result)
            if err != nil {
                log.Fatal(err)
            }

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
