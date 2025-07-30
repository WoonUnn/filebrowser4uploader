package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func ConfigContentSSEHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	for {
		select {
		case msg := <-ConfigContentChan:
			_, err := w.Write([]byte("data: " + msg + "\n\n"))
			if err != nil {
				http.Error(w, "Error writing response", http.StatusInternalServerError)
				return
			}
			flusher.Flush()
		case <-r.Context().Done():
			log.Println("Client disconnected")
			return
		}
	}
}

func ConfigContentHandler(w http.ResponseWriter, r *http.Request) {
	androidCfg := &androidUploadingContent{
		EnableLogUpload:   NowConfigContent.Log,
		EnableImageUpload: NowConfigContent.Image,
		EnableVideoUpload: NowConfigContent.Video,
	}

	msg, err := json.Marshal(androidCfg)
	if err != nil {
		http.Error(w, "Error marshalling config content", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(msg)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}