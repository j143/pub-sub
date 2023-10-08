package main

import (
	"sync"
)

// PubSub represents the pub-sub system.
type PubSub struct {
	mu      sync.Mutex
	topics  map[string][]chan string
	queues  map[string][]string //keys are topic names, values are slices of strings
	history map[string][]string
}

// NewPubSub creates a new PubSub instance.
func NewPubSub() *PubSub {
	return &PubSub{
		topics: make(map[string][]chan string),
	}
}

// Subscribe adds a new subscriber to a topic.
func (ps *PubSub) Subscribe(topic string) chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan string)
	ps.topics[topic] = append(ps.topics[topic], ch)

	return ch
}

// Publish sends a message to all subscribers of a topic.
func (ps *PubSub) Publish(topic string, message string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	subscribers, ok := ps.topics[topic]

	if !ok {
		ps.EnqueueMessage(topic, message)
		ps.PersistMessage(topic, message)
		return
	}

	for _, ch := range subscribers {
		go func(ch chan string) {
			ch <- message
		}(ch)
	}
}


// Unsubscribe removes a subscriber from a topic.
func (ps *PubSub) Unsubscribe(topic string, ch chan string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	subscribers, ok := ps.topics[topic]
	if !ok {
		return
	}

	for i, subscriber := range subscribers {
		if subscriber == ch {
			// Remove the subscriber by replacing it with the last subscriber in the slice
			subscribers[i] = subscribers[len(subscribers)-1]
			ps.topics[topic] = subscribers[:len(subscribers)-1]
			close(ch) // Close the channel to prevent memory leaks
			break
		}
	}
}


func (ps *PubSub) EnqueueMessage(topic string, message string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, ok := ps.queues[topic]; !ok {
		ps.queues[topic] = []string{message}
	} else {
		ps.queues[topic] = append(ps.queues[topic], message)
	}
}


func (ps *PubSub) ProcessQueue(topic string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	messages, ok := ps.queues[topic]
	if !ok {
		return
	}

	for _, message := range messages {
		ps.Publish(topic, message)
	}

	delete(ps.queues, topic)
}

func (ps *PubSub) PersistMessage(topic string, message string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if _, ok := ps.history[topic]; !ok {
		ps.history[topic] = []string{message}
	} else {
		ps.history[topic] = append(ps.history[topic], message)
	}
}

func (ps *PubSub) RetrieveHistory(topic string, ch chan string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	history, ok := ps.history[topic]
	if !ok {
		return
	}

	for _, message := range history {
		ch <- message
	}
}
