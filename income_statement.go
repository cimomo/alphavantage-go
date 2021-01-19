package alphavantage

// IncomeStatementReport defines the response body of the GetIncomeStatement() API.
type IncomeStatementReport struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	TotalRevenue                      string `json:"totalRevenue"`
	TotalOperatingExpense             string `json:"totalOperatingExpense"`
	CostOfRevenue                     string `json:"costOfRevenue"`
	GrossProfit                       string `json:"grossProfit"`
	EBIT                              string `json:"ebit"`
	NetIncome                         string `json:"netIncome"`
	ResearchAndDevelopment            string `json:"researchAndDevelopment"`
	EffectOfAccountingCharges         string `json:"effectOfAccountingCharges"`
	IncomeBeforeTax                   string `json:"incomeBeforeTax"`
	MinorityInterest                  string `json:"minorityInterest"`
	SellingGeneralAdministrative      string `json:"sellingGeneralAdministrative"`
	OtherNonOperatingIncome           string `json:"otherNonOperatingIncome"`
	OperatingIncome                   string `json:"operatingIncome"`
	OtherOperatingExpense             string `json:"otherOperatingExpense"`
	InterestExpense                   string `json:"interestExpense"`
	TaxProvision                      string `json:"taxProvision"`
	InterestIncome                    string `json:"interestIncome"`
	NetInterestIncome                 string `json:"netInterestIncome"`
	ExtraordinaryItems                string `json:"extraordinaryItems"`
	NonRecurring                      string `json:"nonRecurring"`
	OtherItems                        string `json:"otherItems"`
	IncomeTaxExpense                  string `json:"incomeTaxExpense"`
	TotalOtherIncomeExpense           string `json:"totalOtherIncomeExpense"`
	DiscontinuedOperations            string `json:"discontinuedOperations"`
	NetIncomeFromContinuingOperations string `json:"netIncomeFromContinuingOperations"`
	NetIncomeApplicableToCommonShares string `json:"netIncomeApplicableToCommonShares"`
	PreferredStockAndOtherAdjustments string `json:"preferredStockAndOtherAdjustments"`
}
