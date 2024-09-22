package main

import (
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	srv := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(myHandler),
	}

	log.Printf("Starting server on port %s", srv.Addr)
	// TODO: handle SIGINT/SIGTERM and call srv.Shutdown() to gracefully
	// shutdown the server
	log.Fatal(srv.ListenAndServe())
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request from %s for %s", r.RemoteAddr, r.URL.Path)
	w.Write([]byte("Hello, World!"))
}
