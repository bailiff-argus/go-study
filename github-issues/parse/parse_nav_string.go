package parse

import (
	"fmt"
	"strings"
)

/*
Example of in data:
<https://api.github.com/search/issues?q=repo%3Amorhetz%2Fgruvbox&page=1>; rel="prev",
<https://api.github.com/search/issues?q=repo%3Amorhetz%2Fgruvbox&page=3>; rel="next",
<https://api.github.com/search/issues?q=repo%3Amorhetz%2Fgruvbox&page=15>; rel="last"
*/

func ParseNavigation (raw string) (map[string]string) {
    if raw == "" {
        return nil
    }

    dirToPage := make(map[string]string)
    s := strings.Split(raw, ",")
    for _, item := range s {
        page, direction := formNavQueryPair(item)
        dirToPage[direction] = page
    }

    for key, val := range dirToPage {
        fmt.Printf("%s:\t%s\n", key, val)
    }

    fmt.Println()

    return dirToPage
}

func formNavQueryPair (line string) (string, string) {
    s := strings.Split(strings.TrimSpace(line), ";")
    query := extractQuery(s[0])
    nav   := extractNav(s[1])

    return query, nav
}

func extractQuery (line string) (string) {
    query := strings.TrimSpace(line)
    query = strings.Split(query[1:len(query)-1], "issues")[1]
    return query
}

func extractNav (line string) (string) {
    nav := strings.TrimSpace(line)
    nav = nav[5:len(line)-2]
    return nav
}
