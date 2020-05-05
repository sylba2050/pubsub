package pubsub

import (
	"net"
	"sync"
)

var m sync.Mutex

// Registry is information registry
type Registry struct {
	topics map[string][]*Client
}

// Publish pusblish data
func (r *Registry) Publish(topicID string, message Message) {
	r.pub(topicID, message)
}

func (r *Registry) pub(topicID string, message Message) {
	go func() {
		for _, client := range r.topics[topicID] {
			m, err := message.Tobytes()
			if err != nil {
				return
			}
			client.c <- m
		}
	}()
}

// Subscribe subscrib client
func (r *Registry) Subscribe(topicID string, client *Client) {
	m.Lock()
	defer m.Unlock()

	r.topics[topicID] = append(r.topics[topicID], client)
}

// Client has client's data
type Client struct {
	c    chan []byte
	conn net.Conn
}

// NewClient create new client data
func NewClient(conn net.Conn) Client {
	c := make(chan []byte, 0)
	return Client{c: c, conn: conn}
}
