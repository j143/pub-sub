package main

import (
	"fmt"
	"time"
)

func main() {
	ps := NewPubSub()

	chA := ps.Subscribe("topicA")
	chB := ps.Subscribe("topicB")
	chC := ps.Subscribe("topicC") // to test unsubscribe functionality

	go func() {
		for {
			select {
			case msg := <-chA:
				fmt.Println("Received on topicA:", msg)
			case msg := <-chB:
				fmt.Println("Received on topicB:", msg)
			case msg := <-chC:
				fmt.Println("Received on topicC:", msg)
			}
		}
	}()

	ps.Publish("topicA", "Hello, topicA subscribers!")
	ps.Publish("topicB", "Greetings, topicB subscribers!")
	ps.Publish("topicC", "Greetings, topicC subscribers!")

	time.Sleep(time.Second)

	ps.ProcessQueue("topicA")
	ps.ProcessQueue("topicB")
	ps.ProcessQueue("topicC")

	// Now, let's simulate a new subscribe who missed previous messages
	newSub := make(chan string)
    ps.RetrieveHistory("topicB", newSub)
	go func() {
		for {
			select {
			case msg := <-newSub:
				fmt.Println("Received messages from history: ", msg)
			}
		}
	}()

	// Unsubscribe from topicC
	// ps.Unsubscribe("topicC", chC)

	// Publish a message after unsubscribing
	// ps.Publish("topicC", "This message won't be received by topicC subscribers")
	// panic: send on closed channel

	time.Sleep(time.Second)
}
