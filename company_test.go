package alphavantage

import (
	"testing"
)

func TestGetCompanyOverview(t *testing.T) {
	client, err := NewClient(apiKey)
	if err != nil {
		t.Error(err)
	}

	result, err := client.GetCompanyOverview("IBM")
	if err != nil {
		t.Error(err)
	}

	if result.Symbol != "IBM" {
		t.Errorf("Symbol expected: %s", "IBM")
	}
}
