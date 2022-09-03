package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"bufio"
	"os/exec"
	"time"

	"go-study/github-issues/github"
	"go-study/github-issues/nav"
	"go-study/github-issues/parse"
	"go-study/github-issues/issue"
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
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()

        navigator := parse.ParseNavigation(navStr)

        fmt.Printf("%d issues:\n", result.TotalCount)
        for _, item := range result.Items {
            fmt.Printf(
                "#%-5d %v %9.9s %.55s\n",
                item.Number, setTimeTranche(item.CreatedAt), item.User.Login, item.Title,
            )
        }

        fmt.Println()
        fmt.Println("[F]irst    [P]revious    [N]ext    [L]ast")
        fmt.Println("[C]reate   [U]pdate #    [R]ead #  [D]elete #")
        fmt.Println()

        // Processing user input
        fmt.Printf("\nCOMMAND | :")
        input, err := reader.ReadString('\n')
        input = input[:len(input)-1]    // remove delimiter
        input = strings.ToLower(input)

        fmt.Printf("Current input: %s\n", input)

        if err != nil {
            log.Fatal(err)
        }

        if input == "q" { // [q]uit
            break
        } else if input == "c" { // [c]reate

        } else if strings.HasPrefix(input, "r") { 
            issueNumber, err := strconv.Atoi(input[2:])
            if err != nil {
                log.Fatal(err)
            }
            cmd.Run()
            issue.ViewIssue(issueNumber, result)

            fmt.Println("Press ENTER to go back")
            reader.ReadString('\n')

        } else if strings.HasPrefix(input, "u") {

        } else if strings.HasPrefix(input, "d") {

        } else { // if not those actions, then navigate
            q, ok := nav.Navigate(input, navigator)
            if !ok {
                continue
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
