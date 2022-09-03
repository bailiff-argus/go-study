package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func SearchIssues (repo string, auth string, searchRes *IssuesSearchResult) (string, *IssuesSearchResult, error) {
    q := "?q=" + url.QueryEscape("repo:" + repo)
    return SendRequest(q, auth, searchRes)
}

func SendRequest (q string, auth string, searchRes *IssuesSearchResult) (string, *IssuesSearchResult, error) {
    fmt.Println(q)
    fmt.Println(IssuesURL + q)

    client := &http.Client{}
    req, _ := http.NewRequest("GET", IssuesURL + q, nil)
    if auth != "" {
        key := "Authorization"
        value := fmt.Sprintf("Bearer %s", auth)
        req.Header.Set(key, value)
        req.Header.Add("Accept", "application/vnd.github+json")
    }

    resp, err := client.Do(req)
    if err != nil {
        return "", nil, err
    }

    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return "", nil, fmt.Errorf("search query failed: %s", resp.Status)
    }

    if err := json.NewDecoder(resp.Body).Decode(searchRes); err != nil {
        resp.Body.Close()
        return "", nil, err
    }

    if searchRes.TotalCount == 0 {
        return "", searchRes, nil
    }

    meta := resp.Header.Get("link")

    resp.Body.Close()
    return meta, searchRes, nil
}
