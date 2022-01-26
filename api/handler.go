// This is the handler for http requests
package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/canberksinangil/cache-memory/cache"
)

// cacheHandler holds the cache in order to use the cache functions
type cacheHandler struct {
	cache *cache.Cache
}

// NewCacheHandler creates a new cache handler
func NewCacheHandler(c *cache.Cache) *cacheHandler {
	return &cacheHandler{
		cache: c,
	}

}

// HealthCheck is a mechanism to understand if api is working
// It only accepts get requests and response 200
func (ch *cacheHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		createHttpResponse(w, http.StatusOK, nil)
	default:
		createHttpResponse(w, http.StatusMethodNotAllowed, nil)
	}
}

// Cache is a router for 'get', 'set' and 'delete'
// It routes based on http methods
func (ch *cacheHandler) Cache(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ch.get(w, r)
	case http.MethodPost:
		ch.set(w, r)
	case http.MethodDelete:
		ch.delete(w, r)
	default:
		createHttpResponse(w, http.StatusMethodNotAllowed, nil)
	}
}

// Get handles get requests and calls get method from cache
// It only accepts get requests
// key is the required parameters
func (ch *cacheHandler) get(w http.ResponseWriter, r *http.Request) {
	resp := new(response)
	key := r.URL.Query().Get("key")
	if len(key) == 0 {
		resp.Error = noKeyError
		createHttpResponse(w, http.StatusBadRequest, resp)
		return
	}

	value, exists := ch.cache.Get(key)
	if !exists {
		resp.Error = fmt.Sprintf(keyNotFoundError, key)
		createHttpResponse(w, http.StatusOK, resp)
		return
	}
	resp.Value = value
	createHttpResponse(w, http.StatusOK, resp)

}

// Set handles set requests and calls set method from cache
// It only accepts json post requests
// request is binded to setRequest
// key and value are the required parameters
func (ch *cacheHandler) set(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	resp := new(response)
	req := new(setRequest)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Error = err.Error()
		createHttpResponse(w, http.StatusBadRequest, resp)
		return
	}

	if len(req.Key) == 0 {
		resp.Error = noKeyError
		createHttpResponse(w, http.StatusBadRequest, resp)
		return
	}

	if len(req.Value) == 0 {
		resp.Error = noValueError
		createHttpResponse(w, http.StatusBadRequest, resp)
		return
	}

	ch.cache.Set(req.Key, req.Value)
	createHttpResponse(w, http.StatusOK, nil)
}

// Delete handles delete requests and calls delete method from cache
// It only accepts json post requests
// request is binded to deleteRequest
// key is the required parameter
func (ch *cacheHandler) delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	resp := new(response)
	req := new(deleteRequest)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Error = err.Error()
		createHttpResponse(w, http.StatusBadRequest, resp)
		return
	}

	if len(req.Key) == 0 {
		resp.Error = noKeyError
		createHttpResponse(w, http.StatusBadRequest, resp)
		return
	}

	ch.cache.Delete(req.Key)
	createHttpResponse(w, http.StatusOK, nil)
}

// Flush handles flush requests and calls flush method from cache
// It only accepts post requests
func (ch *cacheHandler) Flush(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		if err := ch.cache.Flush(); err != nil {
			resp := new(response)
			resp.Error = err.Error()
			log.Println(err)
			createHttpResponse(w, http.StatusInternalServerError, nil)
		}
		createHttpResponse(w, http.StatusOK, nil)
	default:
		createHttpResponse(w, http.StatusMethodNotAllowed, nil)
	}
}
