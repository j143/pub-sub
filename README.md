

# Concurrent In-Memory Pub-Sub System in Go

This is a simple concurrent in-memory pub-sub system implemented in Go. It allows multiple subscribers to listen to different topics and receive messages concurrently.

## Usage

1. Make sure you have Go installed on your system.

2. Clone the repository:

   ```sh
   git clone https://github.com/j143/pub-sub.git
   ```

3. Navigate to the project directory:

   ```sh
   cd pub-sub
   ```

4. Run the program:

   ```sh
   go run main.go pubsub.go
   ```

## How It Works

The `PubSub` struct handles the subscription and publishing of messages. It uses channels for communication between subscribers and the pub-sub system.

### Creating a PubSub Instance

```go
ps := NewPubSub()
```

### Subscribing to a Topic

```go
chA := ps.Subscribe("topicA")
chB := ps.Subscribe("topicB")
```

### Publishing a Message

```go
ps.Publish("topicA", "Hello, topicA subscribers!")
ps.Publish("topicB", "Greetings, topicB subscribers!")
```

### Receiving Messages

```go
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
```

## Extra functionality

Certainly! Adding complexity to the existing code can make it more robust and versatile. Here are some ideas to enhance the pub-sub system:

- [x] 1. **Unsubscribe Functionality**:
   - Implement a method to allow subscribers to unsubscribe from a topic.

- [ ] 2. **Error Handling**:
   - Add proper error handling for cases like attempting to subscribe to a non-existent topic or handling closed channels.

- [ ] 3. **Non-Blocking Publish**:
   - Implement a non-blocking version of the `Publish` function that sends messages without blocking the publisher.

- [ ] 4. **Message Queues**:
   - Introduce a message queue system that can handle a backlog of messages and ensure they are delivered to subscribers even if they're not actively listening.

- [ ] 5. **Message Persistence**:
   - Add a feature to persist messages so that subscribers can retrieve missed messages when they subscribe to a topic.

- [ ] 6. **Wildcard Subscriptions**:
   - Allow subscribers to use wildcard characters in topic names to subscribe to multiple related topics at once (e.g., subscribing to all topics starting with "weather/").

- [ ] 7. **Rate Limiting**:
   - Implement a mechanism to limit the rate at which messages are published to prevent flooding subscribers with messages.

- [ ] 8. **Timeouts**:
   - Add timeouts for subscribers, so they automatically unsubscribe if they don't receive messages within a certain time frame.

- [ ] 9. **Metrics and Monitoring**:
   - Implement a way to track metrics like message delivery rate, subscriber count, etc., for monitoring and analysis purposes.

- [ ] 10. **Distributed Pub-Sub**:
    - Modify the system to work in a distributed environment with multiple servers, possibly using technologies like Redis for message queuing.

- [ ] 11. **Authentication and Authorization**:
    - Add a layer for authentication and authorization to ensure that only authorized users can publish or subscribe to specific topics.
