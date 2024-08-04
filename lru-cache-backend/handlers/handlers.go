// Defines HTTP handlers to manage cache operations.

package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"lru-cache-backend/cache"
)

type CacheHandler struct {
	cache *cache.LRUCache
}

//cache operation request
type SetRequest struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Expiration int64  `json:"expiration"` // Expiration in seconds
}

func NewCacheHandler(cache *cache.LRUCache) *CacheHandler {
	return &CacheHandler{cache: cache}
}

// GetCacheHandler handles GET requests to retrieve a value from the cache.
func (h *CacheHandler) GetCacheHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, found := h.cache.Get(key)
	if !found {
		http.Error(w, "Key not found or expired", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", value)
}

// SetCacheHandler handles POST requests to add a new key-value pair to the cache.
func (h *CacheHandler) SetCacheHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var req SetRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Expiration <= 0 {
		req.Expiration = 5 // Default expiration is 5 seconds
	}

	h.cache.Set(req.Key, req.Value, req.Expiration)
	w.WriteHeader(http.StatusOK)
}

// DeleteCacheHandler handles DELETE requests to remove a key-value pair from the cache.
func (h *CacheHandler) DeleteCacheHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	h.cache.Delete(key)
	w.WriteHeader(http.StatusOK)
}
