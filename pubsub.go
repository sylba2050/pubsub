package pubsub

import (
	"sync"
)

var m sync.Mutex

// Registry is information registry
type Registry interface {
	Subscribe(string, *Client) error
	Publish(string, Message) error
}

// R has topic list and subscribing client list
type R struct {
	// map[topic_id]map[client_id]*Client
	topics map[string]map[string]*Client
}

// Publish pusblish data
func (r *R) Publish(topicID string, message Message) error {
	return r.pub(topicID, message)
}

func (r *R) pub(topicID string, message Message) error {
	go func() {
		for _, client := range r.topics[topicID] {
			(*client).ReceiveMessage(message)
		}
	}()
	return nil
}

// Subscribe subscrib client
func (r *R) Subscribe(topicID string, client *Client) error {
	m.Lock()
	defer m.Unlock()

	clientID, err := (*client).GetID()
	if err != nil {
		return err
	}

	r.topics[topicID][clientID] = client
	return nil
}
