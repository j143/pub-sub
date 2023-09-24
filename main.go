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
}
