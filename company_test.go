package alphavantage

import (
	"testing"
)

func TestGetCompanyOverview(t *testing.T) {
	client, err := NewClient(apiKey)
	if err != nil {
		t.Error(err)
	}

	company, err := client.GetCompanyOverview("IBM")
	if err != nil {
		t.Error(err)
	}

	if company.Symbol != "IBM" {
		t.Errorf("Symbol expected: %s", "IBM")
	}
}
