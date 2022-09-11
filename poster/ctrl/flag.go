package ctrl

import (
    "flag"
)

func GetFlags() (string, string, string) {
    movieName  := flag.String("name", "", "the name of the movie")
    authToken  := flag.String("auth", "", "the API token for OMDB")
    outputName := flag.String("o", "out", "the name for output file")
    flag.Parse()

    outFileName := *outputName + ".jpg"

    return *movieName, *authToken, outFileName
}
