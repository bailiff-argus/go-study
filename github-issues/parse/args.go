package parse

import (
    "flag"
)

func ReadFlags () (string, string) {
    var authToken, repo string

    authTokenAd := flag.String("a", "", "GitHub authentication token")
    repoAd      := flag.String("r", "", "GitHub repo in AUTHOR/NAME format")

    flag.Parse()

    authToken = *authTokenAd
    repo = "repo:" + *repoAd

    return authToken, repo
}
