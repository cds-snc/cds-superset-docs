package handlers

import (
	"log"
	"net/http"
)

// RobotsHandler serves the robots.txt file to control web crawler behavior
type RobotsHandler struct {
	robotsContent []byte
}

// NewRobotsHandler creates a new handler for robots.txt requests
func NewRobotsHandler() *RobotsHandler {
	return &RobotsHandler{
		robotsContent: []byte("User-agent: *\nDisallow: /"),
	}
}

// ServeHTTP implements the http.Handler interface
func (h *RobotsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=604800")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write(h.robotsContent)
	if err != nil {
		log.Printf("Error writing robots.txt response: %v", err)
	}
}
