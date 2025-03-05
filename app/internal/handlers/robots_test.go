package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNewRobotsHandler verifies that the NewRobotsHandler function
// initializes a RobotsHandler with the correct default content
func TestNewRobotsHandler(t *testing.T) {
	handler := NewRobotsHandler()

	expectedContent := "User-agent: *\nDisallow: /"
	actualContent := string(handler.robotsContent)

	if actualContent != expectedContent {
		t.Errorf("Expected robots content to be %q, got %q", expectedContent, actualContent)
	}
}

// TestRobotsHandlerServeHTTP verifies that the RobotsHandler responds correctly to HTTP requests
func TestRobotsHandlerServeHTTP(t *testing.T) {
	handler := NewRobotsHandler()

	// Create a test HTTP request
	req, err := http.NewRequest("GET", "/robots.txt", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the content type
	expectedContentType := "text/plain; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Handler returned wrong content type: got %v want %v", contentType, expectedContentType)
	}

	// Check the cache control header
	expectedCacheControl := "public, max-age=604800"
	if cacheControl := rr.Header().Get("Cache-Control"); cacheControl != expectedCacheControl {
		t.Errorf("Handler returned wrong cache control header: got %v want %v", cacheControl, expectedCacheControl)
	}

	// Check the response body
	expectedBody := "User-agent: *\nDisallow: /"
	if body := rr.Body.String(); body != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expectedBody)
	}
}

// TestRobotsHandlerWithCustomContent verifies that the RobotsHandler works with custom robots.txt content
func TestRobotsHandlerWithCustomContent(t *testing.T) {
	customContent := "User-agent: Googlebot\nAllow: /\nUser-agent: *\nDisallow: /private/"

	handler := &RobotsHandler{
		robotsContent: []byte(customContent),
	}

	req, err := http.NewRequest("GET", "/robots.txt", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if body := rr.Body.String(); body != customContent {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, customContent)
	}
}

// TestRobotsHandlerDifferentMethods checks that the handler responds the same way
// regardless of HTTP method used
func TestRobotsHandlerDifferentMethods(t *testing.T) {
	methods := []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}
	handler := NewRobotsHandler()

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			req, err := http.NewRequest(method, "/robots.txt", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("%s: Handler returned wrong status code: got %v want %v",
					method, status, http.StatusOK)
			}

			expectedBody := "User-agent: *\nDisallow: /"
			if body := rr.Body.String(); body != expectedBody {
				t.Errorf("%s: Handler returned unexpected body: got %v want %v",
					method, body, expectedBody)
			}
		})
	}
}
