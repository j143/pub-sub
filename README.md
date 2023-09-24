

# Concurrent In-Memory Pub-Sub System in Go

This is a simple concurrent in-memory pub-sub system implemented in Go. It allows multiple subscribers to listen to different topics and receive messages concurrently.

## Usage

1. Make sure you have Go installed on your system.

2. Clone the repository:

   ```sh
   git clone https://github.com/your-username/pubsub-go.git
   ```

3. Navigate to the project directory:

   ```sh
   cd pubsub-go
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

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
