package config

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
)

func TestLoad_LambdaCredentials(t *testing.T) {
	tests := []struct {
		name       string
		status     int
		response   string
		wantErr    string
		missingEnv string
	}{
		{
			name:     "decrypts both parameters regardless of response order",
			status:   http.StatusOK,
			response: `{"Parameters":[{"Name":"/wordpress/password","Value":"ssm-password"},{"Name":"/wordpress/username","Value":"ssm-user"}]}`,
		},
		{
			name:     "reports SSM access failures",
			status:   http.StatusBadRequest,
			response: `{"__type":"AccessDeniedException","message":"Access denied"}`,
			wantErr:  "load WordPress credentials from SSM",
		},
		{
			name:     "reports invalid parameters",
			status:   http.StatusOK,
			response: `{"InvalidParameters":["/wordpress/password"]}`,
			wantErr:  "SSM parameters not found",
		},
		{
			name:     "rejects a missing username",
			status:   http.StatusOK,
			response: `{"Parameters":[{"Name":"/wordpress/password","Value":"ssm-password"}]}`,
			wantErr:  `SSM parameter "/wordpress/username" is missing or empty`,
		},
		{
			name:     "rejects an empty password",
			status:   http.StatusOK,
			response: `{"Parameters":[{"Name":"/wordpress/username","Value":"ssm-user"},{"Name":"/wordpress/password","Value":""}]}`,
			wantErr:  `SSM parameter "/wordpress/password" is missing or empty`,
		},
		{
			name:     "rejects an absent value",
			status:   http.StatusOK,
			response: `{"Parameters":[{"Name":"/wordpress/username"},{"Name":"/wordpress/password","Value":"ssm-password"}]}`,
			wantErr:  `SSM parameter "/wordpress/username" is missing or empty`,
		},
		{
			name:       "requires the username parameter despite legacy credentials",
			missingEnv: "WORDPRESS_USERNAME_PARAM",
			wantErr:    "WORDPRESS_USERNAME_PARAM",
		},
		{
			name:       "requires the password parameter despite legacy credentials",
			missingEnv: "WORDPRESS_PASSWORD_PARAM",
			wantErr:    "WORDPRESS_PASSWORD_PARAM",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for name, value := range map[string]string{
				"PORT":                     "",
				"SECURITY_TXT_CONTENT":     "Security text",
				"SITE_NAME_EN":             "English",
				"SITE_NAME_FR":             "French",
				"WORDPRESS_URL":            "https://example.com",
				"WORDPRESS_MENU_ID_EN":     "1",
				"WORDPRESS_MENU_ID_FR":     "2",
				"WORDPRESS_USERNAME":       "legacy-user",
				"WORDPRESS_PASSWORD":       "legacy-password",
				"WORDPRESS_USERNAME_PARAM": "/wordpress/username",
				"WORDPRESS_PASSWORD_PARAM": "/wordpress/password",
				"AWS_REGION":               "ca-central-1",
				"AWS_ACCESS_KEY_ID":        "test",
				"AWS_SECRET_ACCESS_KEY":    "test",
				"AWS_SESSION_TOKEN":        "",
				"AWS_MAX_ATTEMPTS":         "1",
			} {
				t.Setenv(name, value)
			}
			if tt.missingEnv != "" {
				t.Setenv(tt.missingEnv, "")
			}

			var calls int
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				if got := r.Header.Get("X-Amz-Target"); got != "AmazonSSM.GetParameters" {
					t.Errorf("Unexpected SSM operation: %q", got)
				}
				var request struct {
					Names          []string
					WithDecryption bool
				}
				if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
					t.Errorf("Decode SSM request: %v", err)
				}
				if !request.WithDecryption || !slices.Equal(request.Names, []string{"/wordpress/username", "/wordpress/password"}) {
					t.Errorf("Unexpected SSM request: %+v", request)
				}
				w.Header().Set("Content-Type", "application/x-amz-json-1.1")
				w.WriteHeader(tt.status)
				if _, err := w.Write([]byte(tt.response)); err != nil {
					t.Errorf("Write SSM response: %v", err)
				}
			}))
			defer server.Close()
			t.Setenv("AWS_ENDPOINT_URL_SSM", server.URL)

			cfg, err := Load()
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("Expected error containing %q, got %v", tt.wantErr, err)
				}
				if cfg != nil {
					t.Fatal("Expected no configuration on failure")
				}
			} else {
				if err != nil {
					t.Fatalf("Load failed: %v", err)
				}
				if cfg.WordPressUsername != "ssm-user" || cfg.WordPressPassword != "ssm-password" {
					t.Fatal("Expected credentials from SSM rather than legacy environment variables")
				}
			}
			wantCalls := 1
			if tt.missingEnv != "" {
				wantCalls = 0
			}
			if calls != wantCalls {
				t.Errorf("Expected %d SSM requests, got %d", wantCalls, calls)
			}
		})
	}
}
