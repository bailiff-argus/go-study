package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SearchIssues (terms []string, searchRes *IssuesSearchResult) (string, *IssuesSearchResult, error) {
    q := url.QueryEscape(strings.Join(terms, " "))

    resp, err := http.Get(IssuesURL + "?q=" + q)
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

    meta := resp.Header.Get("link")

    resp.Body.Close()
    return meta, searchRes, nil
}
