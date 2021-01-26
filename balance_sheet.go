package alphavantage

import (
	"encoding/json"

	"github.com/cimomo/alphavantage-go/rest"
)

// BalanceSheetReport defines a report (annual or quarterly) in the response body of the GetBalanceSheet() API.
type BalanceSheetReport struct {
	FiscalDateEnding                string `json:"fiscalDateEnding"`
	ReportedCurrency                string `json:"reportedCurrency"`
	TotalAssets                     string `json:"totalAssets"`
	IntangibleAssets                string `json:"intangibleAssets"`
	EarningAssets                   string `json:"earningAssets"`
	OtherCurrentAssets              string `json:"otherCurrentAssets"`
	TotalLiabilities                string `json:"totalLiabilities"`
	TotalShareholderEquity          string `json:"totalShareholderEquity"`
	DeferredLongTermLiabilities     string `json:"deferredLongTermLiabilities"`
	OtherCurrentLiabilities         string `json:"otherCurrentLiabilities"`
	CommonStock                     string `json:"commonStock"`
	RetainedEarnings                string `json:"retainedEarnings"`
	OtherLiabilities                string `json:"otherLiabilities"`
	Goodwill                        string `json:"goodwill"`
	OtherAssets                     string `json:"otherAssets"`
	Cash                            string `json:"cash"`
	TotalCurrentLiabilities         string `json:"totalCurrentLiabilities"`
	ShortTermDebt                   string `json:"shortTermDebt"`
	CurrentLongTermDebt             string `json:"currentLongTermDebt"`
	OtherShareholderEquity          string `json:"otherShareholderEquity"`
	PropertyPlantEquipment          string `json:"propertyPlantEquipment"`
	TotalCurrentAssets              string `json:"totalCurrentAssets"`
	LongTermInvestments             string `json:"longTermInvestments"`
	NetTangibleAssets               string `json:"netTangibleAssets"`
	ShortTermInvestments            string `json:"shortTermInvestments"`
	NetReceivables                  string `json:"netReceivables"`
	LongTermDebt                    string `json:"longTermDebt"`
	Inventory                       string `json:"inventory"`
	AccountsPayable                 string `json:"accountsPayable"`
	TotalPermanentEquity            string `json:"totalPermanentEquity"`
	AdditionalPaidInCapital         string `json:"additionalPaidInCapital"`
	CommonStockTotalEquity          string `json:"commonStockTotalEquity"`
	PreferredStockTotalEquity       string `json:"preferredStockTotalEquity"`
	RetainedEarningsTotalEquity     string `json:"retainedEarningsTotalEquity"`
	TreasuryStock                   string `json:"treasuryStock"`
	AccumulatedAmortization         string `json:"accumulatedAmortization"`
	OtherNonCurrrentAssets          string `json:"otherNonCurrrentAssets"`
	DeferredLongTermAssetCharges    string `json:"deferredLongTermAssetCharges"`
	TotalNonCurrentAssets           string `json:"totalNonCurrentAssets"`
	CapitalLeaseObligations         string `json:"capitalLeaseObligations"`
	TotalLongTermDebt               string `json:"totalLongTermDebt"`
	OtherNonCurrentLiabilities      string `json:"otherNonCurrentLiabilities"`
	TotalNonCurrentLiabilities      string `json:"totalNonCurrentLiabilities"`
	NegativeGoodwill                string `json:"negativeGoodwill"`
	Warrants                        string `json:"warrants"`
	PreferredStockRedeemable        string `json:"preferredStockRedeemable"`
	CapitalSurplus                  string `json:"capitalSurplus"`
	LiabilitiesAndShareholderEquity string `json:"liabilitiesAndShareholderEquity"`
	CashAndShortTermInvestments     string `json:"cashAndShortTermInvestments"`
	AccumulatedDepreciation         string `json:"accumulatedDepreciation"`
	CommonStockSharesOutstanding    string `json:"commonStockSharesOutstanding"`
}

// BalanceSheet defines the response body of the GetBalanceSheet() API.
type BalanceSheet struct {
	Symbol           string               `json:"symbol"`
	AnnualReports    []BalanceSheetReport `json:"annualReports"`
	QuarterlyReports []BalanceSheetReport `json:"quarterlyReports"`
}

// GetBalanceSheet returns the annual and quarterly balance sheet.
func (client *Client) GetBalanceSheet(symbol string) (*BalanceSheet, error) {
	request := rest.NewRequest("BALANCE_SHEET", symbol, client.APIKey())

	response, err := request.Get(apiServer)
	if err != nil {
		return nil, err
	}

	err = response.CheckStatusCode()
	if err != nil {
		return nil, err
	}

	result, err := parseBalanceSheetResult(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func parseBalanceSheetResult(response *rest.Response) (*BalanceSheet, error) {
	result := BalanceSheet{}

	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
