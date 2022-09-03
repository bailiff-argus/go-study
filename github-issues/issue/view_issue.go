package issue

// Headers:
// Accept: application/vnd.github+json
// Authorization: Bearer <TOKEN>

import (
    "fmt"
    "go-study/github-issues/github"
)

func findIssue (issueNo int, issues *github.IssuesSearchResult) (*github.Issue, error) {
    for _, issue := range issues.Items {
        if issue.Number == issueNo {
            return issue, nil
        }
    }

    return nil, fmt.Errorf("no issue with such number on current page")
}

func ViewIssue (issueNo int, issues *github.IssuesSearchResult) (error) {
    issue, err := findIssue(issueNo, issues)
    if err != nil {
        return err
    }

    fmt.Printf("%s\t%s\n\n%s\n", issue.User.Login, issue.CreatedAt, issue.Body)
    return nil
}
