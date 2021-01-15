package alphavantage

const apiServer = "https://www.alphavantage.co/query"

// Client is the API client for the AlphaVantage.
type Client struct {
	apiKey string
}

// NewClient creates a new AlphaVantage API client.
func NewClient(apiKey string) (*Client, error) {
	return &Client{apiKey: apiKey}, nil
}

// APIKey field.
func (client *Client) APIKey() string {
	return client.apiKey
}
