package pubsub

import "net"

// Registry is information registry
type Registry struct {
	// map[topicID]subscriberID
	topics map[string]string

	// map[subscriberID]subscribers
	subscribers map[string]Subscriber
}

func (r *Registry) Publish(topic string, body Payload) {

}

type Subscriber struct {
	c    chan interface{}
	conn net.Conn
}
