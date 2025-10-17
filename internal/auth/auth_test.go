package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test: No auth header included
	h := http.Header{}
	_, err := GetAPIKey(h)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Test: No auth header included, want: err == %v, got: err == %v", ErrNoAuthHeaderIncluded, err)
	}

	// Test: malformed authorization header - missing value
	h["Authorization"] = []string{}
	_, err = GetAPIKey(h)
	if err == nil {
		t.Errorf("Test: malformed authorization header, want: err == malformed authorization header, got: err == %v", err)
	}

	// Test: malformed authorization header - unexpected value
	h["Authorization"] = []string{"UnknownKey SomeValue"}
	_, err = GetAPIKey(h)
	if err == nil {
		t.Errorf("Test: malformed authorization header, want: err == malformed authorization header, got: err == %v", err)
	}

	// Test: valid authorization header
	h["Authorization"] = []string{"ApiKey somevalue"}
	apiKey, err := GetAPIKey(h)
	if err != nil {
		t.Errorf("Test: valid authorization header, want: err == nil, got err == %v", err)
	}
	if apiKey == "" {
		t.Errorf("Test: valid authorization header, want: apiKey == %v, got: apiKey == %v", apiKey, "")
	}
}
