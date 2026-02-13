package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNewSecurityHandler verifies that the NewSecurityHandler function
// initializes a SecurityHandler with the provided content
func TestNewSecurityHandler(t *testing.T) {
	expectedContent := "Contact: security@example.com\nExpires: 2027-12-31T23:59:59Z"
	handler := NewSecurityHandler(expectedContent)

	actualContent := string(handler.securityContent)

	if actualContent != expectedContent {
		t.Errorf("Expected security content to be %q, got %q", expectedContent, actualContent)
	}
}

// TestSecurityHandlerServeHTTP verifies that the SecurityHandler responds correctly to HTTP requests
func TestSecurityHandlerServeHTTP(t *testing.T) {
	content := "Contact: security@example.com\nExpires: 2027-12-31T23:59:59Z"
	handler := NewSecurityHandler(content)

	// Create a test HTTP request
	req, err := http.NewRequest("GET", "/.well-known/security.txt", nil)
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
	if body := rr.Body.String(); body != content {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, content)
	}
}

// TestSecurityHandlerWithEmptyContent verifies that the SecurityHandler works with empty content
func TestSecurityHandlerWithEmptyContent(t *testing.T) {
	handler := NewSecurityHandler("")

	req, err := http.NewRequest("GET", "/.well-known/security.txt", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if body := rr.Body.String(); body != "" {
		t.Errorf("Handler returned unexpected body: got %v want empty string", body)
	}
}

// TestSecurityHandlerWithMultilineContent verifies that the SecurityHandler works with multi-line security.txt content
func TestSecurityHandlerWithMultilineContent(t *testing.T) {
	customContent := `Contact: mailto:security@example.com
Contact: https://example.com/security
Encryption: https://example.com/pgp-key.txt
Acknowledgments: https://example.com/hall-of-fame
Preferred-Languages: en, fr
Canonical: https://example.com/.well-known/security.txt
Policy: https://example.com/security-policy
Expires: 2027-12-31T23:59:59Z`

	handler := NewSecurityHandler(customContent)

	req, err := http.NewRequest("GET", "/.well-known/security.txt", nil)
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

// TestSecurityHandlerDifferentMethods checks that the handler responds the same way
// regardless of HTTP method used
func TestSecurityHandlerDifferentMethods(t *testing.T) {
	methods := []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}
	content := "Contact: security@example.com\nExpires: 2027-12-31T23:59:59Z"
	handler := NewSecurityHandler(content)

	for _, method := range methods {
		t.Run(method, func(t *testing.T) {
			req, err := http.NewRequest(method, "/.well-known/security.txt", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("%s: Handler returned wrong status code: got %v want %v",
					method, status, http.StatusOK)
			}

			if body := rr.Body.String(); body != content {
				t.Errorf("%s: Handler returned unexpected body: got %v want %v",
					method, body, content)
			}
		})
	}
}
