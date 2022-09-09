package ctrl

import (
    "flag"
)

func ReadFlags() (bool, string) {
    forceRebuild := flag.Bool("rebuild", false, "rebuild offline index")
    searchTerm   := flag.String("search", "", "find comics matching description")
    flag.Parse()
    return *forceRebuild, *searchTerm
}
