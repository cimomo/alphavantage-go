package alphavantage

import (
	"testing"
)

// Replace this with a valid API key before running the tests.
const apiKey = "itisstrangethatthiskeyworks"

func TestNewClient(t *testing.T) {
	client, err := NewClient(apiKey)
	if err != nil {
		t.Error(err)
	}

	if client.APIKey() != apiKey {
		t.Errorf("APIKey expected: %s", apiKey)
	}
}
