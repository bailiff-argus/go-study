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
    logFile, err := os.Create("log.log")
    if err != nil {
        log.Fatal(err)
    }
    defer logFile.Close()

    log.SetOutput(logFile)

    token, repo := parse.ReadFlags()

    result := new(github.IssuesSearchResult) // contains current issues page
    var navStr string

    // for auth need flags

    navStr, result, err = github.SearchIssues(repo, token, result)
    if err != nil {
        log.Fatal(err)
    }

    navigator := parse.ParseNavigation(navStr)
    reader := bufio.NewReader(os.Stdin)

    // Mainloop
    for {
        clear()
        fmt.Println(os.Getenv("EDITOR"))

        parse.DisplayResult(result)
        showInterface()

        // Processing user input
        input, err := parse.ReceiveInput(reader)
        if err != nil {
            log.Fatal(err)
        }

        if input == "q" { // [q]uit
            break

        } else if strings.HasPrefix(input, "c") { // [c]reate
            err := issue.CreateIssue(input[2:], repo, token)
            if err != nil {
                log.Printf("github-issues: %v\n", err)
                continue
            }

        } else if strings.HasPrefix(input, "r") { // [r]ead
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

        } else if strings.HasPrefix(input, "u") { // [u]pdate
            issueNumber, err := strconv.Atoi(input[2:])
            if err != nil {
                log.Printf("github-issues: %v\n", err)
                continue
            }

            err = issue.UpdateIssue(issueNumber, result, token)
            if err != nil {
                log.Printf("github-issues: %v\n", err)
                continue
            }

        } else if strings.HasPrefix(input, "d") { // [d]elete

        } else { // if not those actions, then attempt to navigate
            q, ok := nav.Navigate(input, navigator)
            if !ok {
                continue
            }

            navStr, result, err = github.SendRequest(q, token, result)
            if err != nil {
                log.Fatal(err)
            }

            navigator = parse.ParseNavigation(navStr)
        }
    }
}

func showInterface() {
    fmt.Println( "\n[F]irst            [P]revious        [N]ext         [L]ast")
    fmt.Println(   "[C]reate <TITLE>   [U]pdate #        [R]ead #       [D]elete #")
    fmt.Println(   "[Q]uit")
    fmt.Printf("\n\nCOMMAND | :")
}

func clear () {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}
