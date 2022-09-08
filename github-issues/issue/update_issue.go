package issue

import (
	"fmt"
	"net/http"

	"go-study/github-issues/github"
	"go-study/github-issues/parse"

    "time"
)

func UpdateIssue (issueNo int, issues *github.IssuesSearchResult, auth string) error {
    if auth == "" {
        return fmt.Errorf("cannot update issue without authorization")
    }

    issue, err := findIssue(issueNo, issues)
    if err != nil {
        return err
    }

    originalTitle := issue.Title
    originalBody  := issue.Body

    newTitle := parse.GetInputFromEditor(originalTitle)
    newBody  := parse.GetInputFromEditor(originalBody)

    body, err := parse.FormRequestBody(newTitle, newBody)
    fmt.Println(body)
    if err != nil {
        return err
    }

    client := &http.Client{}
    link := issue.URL

    req, _ := http.NewRequest("PATCH", link, body)
    req.Header.Add("Accept", "application/vnd.github+json")

    authKey := "Authorization"
    authValue := fmt.Sprintf("Bearer %s", auth)
    req.Header.Set(authKey, authValue)

    fmt.Println(req)
    fmt.Println(req.Body)

    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    fmt.Println(resp.Status)
    fmt.Println(resp.Body)
    time.Sleep(60 * time.Second)

    resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("issue update failed: %s", resp.Status)
    }

    return nil
}
