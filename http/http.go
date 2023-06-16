// Package http provides a set of HTTP Cloud Functions samples.
package http

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	// Register an HTTP function with the Functions Framework
	functions.HTTP("HelloHTTPMethod", HelloHTTPMethod)
}

// HelloHTTPMethod is an HTTP Cloud function.
// It uses the request method to differentiate the response.
func HelloHTTPMethod(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "Hello World!")
	case http.MethodPut:
		http.Error(w, "403 - Forbidden", http.StatusForbidden)
	default:
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
