package alphavantage

// CashFlowReport defines a report (annual or quarterly) in the response body of the GetCashFlow() API.
type CashFlowReport struct {
	FiscalDateEnding               string `json:"fiscalDateEnding"`
	ReportedCurrency               string `json:"reportedCurrency"`
	Investments                    string `json:"investments"`
	ChangeInLiabilities            string `json:"changeInLiabilities"`
	CashflowFromInvestment         string `json:"cashflowFromInvestment"`
	OtherCashflowFromInvestment    string `json:"otherCashflowFromInvestment"`
	NetBorrowings                  string `json:"netBorrowings"`
	CashflowFromFinancing          string `json:"cashflowFromFinancing"`
	OtherCashflowFromFinancing     string `json:"otherCashflowFromFinancing"`
	ChangeInOperatingActivities    string `json:"changeInOperatingActivities"`
	NetIncome                      string `json:"netIncome"`
	ChangeInCash                   string `json:"changeInCash"`
	OperatingCashflow              string `json:"operatingCashflow"`
	OtherOperatingCashflow         string `json:"otherOperatingCashflow"`
	Depreciation                   string `json:"depreciation"`
	DividendPayout                 string `json:"dividendPayout"`
	StockSaleAndPurchase           string `json:"stockSaleAndPurchase"`
	ChangeInInventory              string `json:"changeInInventory"`
	ChangeInAccountReceivables     string `json:"changeInAccountReceivables"`
	ChangeInNetIncome              string `json:"changeInNetIncome"`
	CapitalExpenditures            string `json:"capitalExpenditures"`
	ChangeInReceivables            string `json:"changeInReceivables"`
	ChangeInExchangeRate           string `json:"changeInExchangeRate"`
	ChangeInCashAndCashEquivalents string `json:"changeInCashAndCashEquivalents"`
}

// CashFlow defines the response body of the GetCashFlow() API.
type CashFlow struct {
	Symbol           string           `json:"symbol"`
	AnnualReports    []CashFlowReport `json:"annualReports"`
	QuarterlyReports []CashFlowReport `json:"quarterlyReports"`
}
