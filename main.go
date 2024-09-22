package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
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

var (
	pageHits     int
	pageHitsLock sync.Mutex
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	pageHitsLock.Lock()
	defer pageHitsLock.Unlock()
	pageHits++

	log.Printf("Request from %s for %s", r.RemoteAddr, r.URL.Path)
	log.Printf("Page hits: %d", pageHits)
	w.Write([]byte(fmt.Sprintf("Hello, World!\nPage hit counter: %d\n", pageHits)))
}
