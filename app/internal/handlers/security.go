package handlers

import (
	"log"
	"net/http"
)

// SecurityHandler serves the security.txt file at /.well-known/security.txt
type SecurityHandler struct {
	securityContent []byte
}

// NewSecurityHandler creates a new handler for security.txt requests
// The content parameter should be populated from an environment variable
func NewSecurityHandler(content string) *SecurityHandler {
	return &SecurityHandler{
		securityContent: []byte(content),
	}
}

// ServeHTTP implements the http.Handler interface
func (h *SecurityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=604800")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write(h.securityContent)
	if err != nil {
		log.Printf("Error writing security.txt response: %v", err)
	}
}
