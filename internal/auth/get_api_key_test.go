package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyReturnsErrorWhenAuthorizationHeaderMissing(t *testing.T) {
	_, err := GetAPIKey(http.Header{})
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
}

func TestGetAPIKeyReturnsKeyWhenAuthorizationHeaderValid(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey secret-key")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if apiKey != "secret-key" {
		t.Fatalf("expected %q, got %q", "secret-key", apiKey)
	}
}
