package main

import "fmt"

type Message struct {
	From string
	Body string
}

func broadcaster(in <-chan Message, out chan<- Message) {
	for message := range in {
		out <- Message{
			From: "server",
			Body: "echo: " + message.Body,
		}
	}
	close(out)
}

func main() {
	fmt.Println("Lesson 41: WebSocket Basics")

	clientToServer := make(chan Message)
	serverToClient := make(chan Message)

	go broadcaster(clientToServer, serverToClient)

	go func() {
		clientToServer <- Message{From: "client", Body: "hello websocket shape"}
		close(clientToServer)
	}()

	for message := range serverToClient {
		fmt.Printf("received from %s -> %s\n", message.From, message.Body)
	}

	fmt.Println("This models message flow before adding a real WebSocket library.")
}
