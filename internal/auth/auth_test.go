package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey")

	got, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("Gettings ApiKey failed: Error: %v", err)
	}

	if got == "" {
		t.Fatalf("Return Value was empty so no string returned")
	}
}
