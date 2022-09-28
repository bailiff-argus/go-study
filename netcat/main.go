package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
    conn, err := net.Dial("tcp", os.Args[1])
    if err != nil { log.Fatal(err) }

    go func() {
        io.Copy(os.Stdout, conn)
        log.Println("done")
    }()

    mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
    _, err := io.Copy(dst, src)
    if err != nil { log.Fatal(err) }
}
