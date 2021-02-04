package alphavantage

import (
	"testing"
)

func TestGetCashFlow(t *testing.T) {
	client, err := NewClient(apiKey)
	if err != nil {
		t.Error(err)
	}

	result, err := client.GetCashFlow("IBM")
	if err != nil {
		t.Error(err)
	}

	if result.Symbol != "IBM" {
		t.Errorf("Symbol expected: %s", "IBM")
	}

	t.Logf("Number of annual reports: %d", len(result.AnnualReports))
	t.Logf("Number of quarterly reports: %d", len(result.QuarterlyReports))
}
