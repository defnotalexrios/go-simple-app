package v1

import (
	"fmt"
	"io"
	"net/http"
)

// HelloHandler is an HTTP handler that
// prints a greeting to the user.
type HelloHandler struct {
}

// NewHelloHandler builds a new HelloHandler.
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (*HelloHandler) Pattern() string {
	return "/hello"
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
