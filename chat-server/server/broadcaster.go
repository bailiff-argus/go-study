package server

func Broadcaster() {
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
