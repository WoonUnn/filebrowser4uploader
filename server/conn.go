package server

import (
	"log"
	"net/http"
)

func SetupServer() error {
	http.HandleFunc("/image", CORSMiddleWare(imageUploadHandler))
	http.HandleFunc("/log", CORSMiddleWare(logUploadHandler))
	http.HandleFunc("/video", CORSMiddleWare(frameHandler))
	
	http.HandleFunc("/config/content", CORSMiddleWare(ConfigContentHandler))
	http.HandleFunc("/config/content/sse", CORSMiddleWare(ConfigContentSSEHandler))

	log.Printf("Uploading server is running on addr: %s\n", ADDR)
	if err := http.ListenAndServe(ADDR, nil); err != nil {
		return err
	}
	return nil
}
