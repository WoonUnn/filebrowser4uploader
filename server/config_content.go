package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pebbe/zmq4"
)

func ConfigContentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}


	subscriber, _ := zmq4.NewSocket(zmq4.SUB)
	defer subscriber.Close()
	subscriber.Connect("tcp://localhost:5555") // Adjust the address as needed
	subscriber.SetSubscribe("config_content")

	for {
		msg, _ := subscriber.Recv(0)
		log.Println("Received message:", msg)
		_, err := fmt.Fprintf(w, msg)
		if err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
			return
		}
		flusher.Flush()
	}
}