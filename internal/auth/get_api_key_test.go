// generate a mock for the GetAPIKey function
package auth

import (
	"net/http"
	"testing"
)

// create a test function for GetAPIKey that returns the api key from the header

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		shouldError   bool
		expectedError error
	}{
		{
			name: "valid API key",
			headers: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "ApiKey test-api-key")
				return h
			}(),
			expectedKey: "test-api-key",
			shouldError: false,
		},
		{
			name:          "missing authorization header",
			headers:       http.Header{},
			expectedKey:   "",
			shouldError:   true,
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed authorization header - no space",
			headers: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "InvalidFormat")
				return h
			}(),
			expectedKey: "",
			shouldError: true,
		},
		{
			name: "malformed authorization header - wrong prefix",
			headers: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "Bearer some-token")
				return h
			}(),
			expectedKey: "",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)
			if tt.shouldError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if apiKey != tt.expectedKey {
					t.Errorf("got %q, want %q", apiKey, tt.expectedKey)
				}
			}
		})
	}
}
