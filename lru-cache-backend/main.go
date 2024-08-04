package main

import (
	"log"
	"net/http"
	"time"
	"lru-cache-backend/cache"
	"lru-cache-backend/handlers"
)

func main() {
	// Create a new LRUCache with a capacity of 10 items
	cache := cache.NewLRUCache(10)

	// Initialize handlers
	cacheHandler := handlers.NewCacheHandler(cache)

	// Define routes
	http.HandleFunc("/api/cache", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			cacheHandler.GetCacheHandler(w, r)
		case http.MethodPost:
			cacheHandler.SetCacheHandler(w, r)
		case http.MethodDelete:

			cacheHandler.DeleteCacheHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	srv := &http.Server{
		Addr: ":8080",
		Handler: nil,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(srv.ListenAndServe())
}
