package pubsub

import (
	"net"
)

type Client interface {
	GetID() (string, error)
	ReceiveMessage(Message) error
	Listen() error
}

// C has client's data
type C struct {
	id             string
	messageChannel chan []byte
	conn           net.Conn
}

func (c C) GetID() (string, error) {
	return c.id, nil
}

func (c *C) ReceiveMessage(m Message) error {
	message, err := m.ToBytes()
	if err != nil {
		return err
	}
	c.messageChannel <- message
	return nil
}

func (c *C) Listen() error {
	go func() {
		for message := range c.messageChannel {
			c.conn.Write(message)
		}
	}()
	return nil
}

// NewClient create new client data
func NewClient(conn net.Conn) C {
	c := make(chan []byte, 0)
	return C{messageChannel: c, conn: conn}
}
