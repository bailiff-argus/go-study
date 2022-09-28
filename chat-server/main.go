package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
    listener, err := net.Listen("tcp", "localhost:8000")
    if err != nil { log.Fatal(err) }

    go broadcaster()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println(err)
            continue
        }

        go handleConn(conn)
    }
}

type client struct {
    receiving chan<- string
    name      string
}

var (
    entering = make(chan client)
    leaving  = make(chan client)
    messages = make(chan string)
)

func broadcaster() {
    clients := make(map[client]bool)
    for {
        select {
        case msg := <-messages:
            for cli := range clients {
                cli.receiving <- msg
            }

        case cli := <-entering:
            clients[cli] = true
            cli.receiving <- "Currently in chat:"
            for client := range clients {
                if client != cli { cli.receiving <- client.name }
            }

        case cli := <-leaving:
            delete(clients, cli)
            close(cli.receiving)
        }
    }
}

func handleConn(conn net.Conn) {
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
