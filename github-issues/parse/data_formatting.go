package parse

import (
    "time"
)

func SetTimeTranche (t time.Time) string {
    now := time.Now()
    monthAgo := now.AddDate(0, -1, 0)
    yearAgo := now.AddDate(-1, 0, 0)

    if t.After(monthAgo) {
        return "Less than a month old"
    } else if t.After(yearAgo) {
        return "Less than a year old "
    } else {
        return "Older than a year    "
    }
}
