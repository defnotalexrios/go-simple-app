package v1

import (
	"io"
	"net/http"
)

// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct {
}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//logger := telemetry.Logger(r.Context())
	if _, err := io.Copy(w, r.Body); err != nil {
		//logger.Warn("Failed to handle request", zap.Error(err))
	}
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}
