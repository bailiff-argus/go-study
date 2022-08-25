package nav

var keymaps = map[string]string{
    "n": "next",
    "p": "prev",
    "f": "first",
    "l": "last",
}

func Navigate(cmd string, navigator map[string]string) (string, bool) {
    var direction, q string
    var ok bool

    direction, ok = keymaps[cmd]
    if !ok {
        return "", false
    }

    q, ok = navigator[direction]
    if !ok {
        return "", false
    }


    if cmd == "n" { // [n]ext
        q, ok = navigator["next"]
        if !ok {
            return "", false
        }
    } else if cmd == "p" { // [p]revious
        q, ok = navigator["prev"]
        if !ok {
            return "", false
        }
    } else if cmd == "f" { // [f]irst
        q, ok = navigator["first"]
        if !ok {
            return "", false
        }
    } else if cmd == "l" { // [l]ast
        q, ok = navigator["last"]
        if !ok {
            return "", false
        }
    }

    return q, true
}
