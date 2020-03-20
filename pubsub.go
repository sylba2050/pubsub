package pubsub

import "net"

// Registry is information registry
type Registry struct {
	// map[topicID]clientID
	topics map[string]string

	// map[clientID]subscribers
	clients map[string]Client
}

func (r *Registry) Publish(topicID string, payload Payload) {
}

func (r *Registry) Subscribe(topicID, subscriberID string) {
}

type Client struct {
	id   string
	c    chan interface{}
	conn net.Conn
}

type NewClient struct{}
