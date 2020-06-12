package pubsub

import (
	"net"
)

type Client interface {
	GetID() (string, error)
	ReceiveMessage(Message) error
	Listen() error
}

// client has client's data
type client struct {
	id             string
	messageChannel chan []byte
	conn           net.Conn
}

func (c client) GetID() (string, error) {
	return c.id, nil
}

func (c *client) ReceiveMessage(m Message) error {
	message, err := m.ToBytes()
	if err != nil {
		return err
	}
	c.messageChannel <- message
	return nil
}

func (c *client) Listen() error {
	go func() {
		for message := range c.messageChannel {
			c.conn.Write(message)
		}
	}()
	return nil
}

// NewClient create new client data
func NewClient(conn net.Conn) client {
	c := make(chan []byte, 0)
	return client{messageChannel: c, conn: conn}
}
