package search

func DepthFirst(f func(item string) []string, worklist []string) {
    seen := make(map[string]bool)
    for len(worklist) > 0 {
        lastIndex := len(worklist) - 1
        item := worklist[lastIndex]
        worklist = worklist[0:lastIndex]

        if !seen[item] {
            seen[item] = true
            worklist = append(worklist, f(item)...)
        }
    }
}
