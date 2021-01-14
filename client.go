package alphavantage

// Client is the API client for the AlphaVantage.
type Client struct {
	apiKey string
}

// NewClient creates a new AlphaVantage API client.
func NewClient(apiKey string) (*Client, error) {
	return &Client{apiKey: apiKey}, nil
}
