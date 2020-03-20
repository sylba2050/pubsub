package pubsub

import "net"

// Registry is information registry
type Registry struct {
	topics map[string][]*Client
}

func (r *Registry) Publish(topicID string, message []byte) error {
	payload := NewPublishPayload(message)
	err := r.pub(topicID, payload)
	return err
}

func (r *Registry) pub(topicID string, payload Payload) error {

	for _,  range r.topics[topicID] {
		r.topics[topicID].c <- payload
	}
}

func (r *Registry) Subscribe(topicID, subscriberID string) {
}

type Client struct {
	c    chan interface{}
	conn net.Conn
}

func NewClient(conn net.Conn) Client {
	c := make(chan []byte, 0)
	return Client{c: c, conn: conn}
}
