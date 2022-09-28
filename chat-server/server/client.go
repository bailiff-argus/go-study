package server

type client struct {
    receiving chan<- string
    name      string
}
