package alphavantage

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient("demo")
	if err != nil {
		t.Error(err)
	}
}
