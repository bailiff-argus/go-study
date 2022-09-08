package issue

import (
    "fmt"
    "net/http"

    "go-study/github-issues/parse"
)

var baseURL string = "https://api.github.com/repos/"

func CreateIssue (title string, repo string, auth string) error {
    if auth == "" {
        return fmt.Errorf("cannot create issue without authorization")
    }


    client := &http.Client{}
    link := formQueryLink(repo)

    text := parse.GetInputFromEditor("")
    body, err := parse.FormRequestBody(title, text)
    if err != nil {
        return err
    }

    req, _ := http.NewRequest("POST", link, body)

    req.Header.Add("Accept", "application/vnd.github+json")

    authKey := "Authorization"
    authValue := fmt.Sprintf("Bearer %s", auth)
    req.Header.Set(authKey, authValue)

    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        return fmt.Errorf("issue creation failed: %s", resp.Status)
    }


    return nil
}

func formQueryLink (repo string) string {
    return fmt.Sprintf("%s%s/issues", baseURL, repo)
}
