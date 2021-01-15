package alphavantage

import (
	"encoding/json"

	"github.com/cimomo/alphavantage-go/rest"
)

// Company defines the response body for the company overview API.
type Company struct {
	Symbol      string `json:"Symbol"`
	AssetType   string `json:"AssetType"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// GetCompanyOverview returns the company information, financial rations, and other key metrics for the equity specified.
func (client *Client) GetCompanyOverview(symbol string) (*Company, error) {
	request := rest.NewRequest("OVERVIEW", symbol, client.APIKey())

	response, err := request.Get(apiServer)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseCompanyOverviewResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func parseCompanyOverviewResult(response *rest.Response) (*Company, error) {
	result := Company{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
