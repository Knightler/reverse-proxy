package main

import (
	"log"
	"net/http"
)

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.
}

func main() {
	log.Println("Starting reverse proxy on :8080")
	http.ListenAndServe(":8080", nil)
}
