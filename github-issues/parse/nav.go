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
    dirToPage := make(map[string]string)
    s := strings.Split(raw, ",")
    for _, item := range s {
        page, direction := formNavLinkPair(item)
        dirToPage[direction] = page
    }

    for key, val := range dirToPage {
        fmt.Printf("%s:\t%s\n", key, val)
    }

    fmt.Println()

    return dirToPage
}

func formNavLinkPair (line string) (string, string) {
    s := strings.Split(strings.TrimSpace(line), ";")
    link := extractLink(s[0])
    nav  := extractNav(s[1])

    return link, nav
}

func extractLink (line string) (string) {
    link := strings.TrimSpace(line)
    link = link[1:len(link)-1]
    return link
}

func extractNav (line string) (string) {
    nav := strings.TrimSpace(line)
    nav = nav[5:len(line)-2]
    return nav
}
