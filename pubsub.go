package pubsub

import (
	"sync"
)

var m sync.Mutex

// Registry is information registry
type Registry interface {
	Subscribe(string, *Client)
	Publish(string, Message)
}

// R has topic list and subscribing client list
type R struct {
	// map[topic_id]map[client_id]*Client
	topics map[string]map[string]*Client
}

// Publish pusblish data
func (r *R) Publish(topicID string, message Message) {
	r.pub(topicID, message)
}

func (r *R) pub(topicID string, message Message) {
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
func (r *R) Subscribe(topicID string, client *Client) {
	m.Lock()
	defer m.Unlock()

	r.topics[topicID][client.id] = client
}

