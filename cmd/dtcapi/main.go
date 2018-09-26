package main

// API server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about/", about)

	fmt.Println("Listinging on port 8081")
	http.ListenAndServe(":8081", nil)
}

type Message struct {
	Text string
}

func about(w http.ResponseWriter, r *http.Request) {
	m := Message{"Welcome to the the distributed_transcoder API, v0.1"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Write(b)
}
