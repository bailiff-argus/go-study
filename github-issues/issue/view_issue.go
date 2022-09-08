package issue

// Headers:
// Accept: application/vnd.github+json
// Authorization: Bearer <TOKEN>

import (
	"fmt"
	"time"

	"go-study/github-issues/github"
	"go-study/github-issues/parse"
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

    text := fmt.Sprintf(
        "%s\t%s\n%s\n\n%s",
        issue.User.Login, issue.CreatedAt.Format(time.RFC822), issue.Title, issue.Body,
    )
    err = parse.ShowInPager(text)
    return err
}
