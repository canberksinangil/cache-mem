// This is a helper for handling http request or server side logging
package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// error messages for the http responses
var (
	noKeyError       = "The 'key' is required."
	keyNotFoundError = "The key '%s' could not be found."
	noValueError     = "The 'value' is required."
)

// response provides flexibility and fields to create http responses
type response struct {
	Error  string `json:"error,omitempty"`
	Result string `json:"result,omitempty"`
	Value  string `json:"value,omitempty"`
}

// setRequest holds to required variables to write key value pairs to the cache
type setRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// deleteRequest holds to required variable to delete a key from the cache
type deleteRequest struct {
	Key string `json:"key"`
}

// createHttpResponse eases creating http responses
func createHttpResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	if response == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// ServerLogger logs http requests
func ServerLogger(handler http.Handler) http.Handler {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
		handler.ServeHTTP(w, r)
	})
}
