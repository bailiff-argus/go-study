package ctrl

import (
    "flag"
)

func ReadFlags() bool {
    forceRebuild := flag.Bool("rebuild", false, "rebuild offline index")
    flag.Parse()
    return *forceRebuild
}
