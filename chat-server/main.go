package main

import (
	"log"
	"net"
    "go-study/chat-server/server"
)

func main() {
    listener, err := net.Listen("tcp", "localhost:8000")
    if err != nil { log.Fatal(err) }

    go server.Broadcaster()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println(err)
            continue
        }

        go server.HandleConn(conn)
    }
}
