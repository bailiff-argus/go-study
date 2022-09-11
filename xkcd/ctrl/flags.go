package ctrl

import (
    "flag"
)

func ReadFlags() (bool, bool, string) {
    forceRebuild := flag.Bool("rebuild", false, "rebuild offline index")
    noUpdate     := flag.Bool("noupdate", false, "use existing online index")
    searchTerm   := flag.String("search", "", "find comics matching description")
    flag.Parse()
    return *forceRebuild, *noUpdate, *searchTerm
}
