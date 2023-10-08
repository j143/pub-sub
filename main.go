package main

import (
	"fmt"
	"time"
)

func main() {
	ps := NewPubSub()

	chA := ps.Subscribe("topicA")
	chB := ps.Subscribe("topicB")

	go func() {
		for {
			select {
			case msg := <-chA:
				fmt.Println("Received on topicA:", msg)
			case msg := <-chB:
				fmt.Println("Received on topicB:", msg)
			}
		}
	}()

	ps.Publish("topicA", "Hello, topicA subscribers!")
	ps.Publish("topicB", "Greetings, topicB subscribers!")

	time.Sleep(time.Second)

	// Unsubscribe from topicA
	// ps.Unsubscribe("topicA", chA)

	ps.ProcessQueue("topicA")
	ps.ProcessQueue("topicB")

	// Publish a message after unsubscribing
	// ps.Publish("topicA", "This message won't be received by topicA subscribers")
	// panic: send on closed channel

// goroutine 20 [running]:
// main.(*PubSub).Publish.func1(0x0?)
//         D:/repo/pub-sub/pubsub.go:44 +0x30
// created by main.(*PubSub).Publish
//         D:/repo/pub-sub/pubsub.go:43 +0x105
// exit status 2

	// Check if the channel is still open before publishing
	select {
	case chA <- "This message won't be received by topicA subscribers":
		fmt.Println("Message sent successfully")
	default:
		fmt.Println("Failed to send message to closed channel")
	}

	time.Sleep(time.Second)
}
