package server

import (
    "fmt"
    "net"
    "time"
    "bufio"
)

var (
    entering = make(chan client)
    leaving  = make(chan client)
    messages = make(chan string)
)

func HandleConn(conn net.Conn) {
    var cli client

    ch := make(chan string, 20)
    go clientWriter(conn, ch)

    timer := time.NewTimer(5 * time.Minute)
    go func() {
        <-timer.C
        conn.Close()
    }()

    cli.receiving = ch
    cli.receiving <- "Enter your name:"
    input := bufio.NewScanner(conn)
    input.Scan()
    cli.name = input.Text()

    cli.receiving <- "You are " + cli.name
    messages <- cli.name + " has arrived"
    entering <- cli

    for input.Scan() {
        timer.Reset(5 * time.Minute)
        messages <- cli.name + ": " + input.Text()
    }

    leaving <- cli
    messages <- cli.name + " has left"
    conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
    for msg := range ch { fmt.Fprintln(conn, msg) }
}
