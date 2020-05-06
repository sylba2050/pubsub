package pubsub

import (
	"net"
)

// Client has client's data
type Client struct {
	id   string
	c    chan []byte
	conn net.Conn
}

// NewClient create new client data
func NewClient(conn net.Conn) Client {
	c := make(chan []byte, 0)
	return Client{c: c, conn: conn}
}
