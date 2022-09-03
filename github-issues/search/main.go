package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"bufio"
	"os/exec"

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
        clear()
        navigator := parse.ParseNavigation(navStr)

        parse.DisplayResult(result)
        showInterface()

        // Processing user input
        input, err := parse.ReceiveInput(reader)
        if err != nil {
            log.Fatal(err)
        }

        if input == "q" { // [q]uit
            break
        } else if input == "c" { // [c]reate

        } else if strings.HasPrefix(input, "r") { 
            clear()

            issueNumber, err := strconv.Atoi(input[2:])
            if err != nil {
                log.Printf("github-issues: %v\n", err)
                continue
            }

            err = issue.ViewIssue(issueNumber, result)
            if err != nil {
                log.Printf("github-issues: %v\n", err)
                continue
            }
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

func showInterface() {
    fmt.Println("\n[F]irst    [P]revious    [N]ext    [L]ast")
    fmt.Println("[C]reate   [U]pdate #    [R]ead #  [D]elete #")
    fmt.Printf("\n\nCOMMAND | :")
}

func clear () {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}
