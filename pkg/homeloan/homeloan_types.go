package homeloan

// AnnualDetail struct defines interesting quants of the EMIs summed up annually
type AnnualDetail struct {

	// AnnualEMIAmount is the total EMI amount paid during AnnualDetail.Year
	AnnualEMIAmount int `json:"annual_emi_amount"`

	// AnnualInterest is the amount from AnnualEMIAmount which is the
	// interest paid during the Year
	AnnualInterest int `json:"annual_interest"`

	// AnuualTaxSaved is the tax saved considering tax benefits u/s 24
	// where borrower can claim exemption upto 2L on interest paid
	// during the financial year
	AnuualTaxSaved int `json:"annual_tax_saved"`

	// MonthlyEMIs is the array of EMIs for the financial year AnnualDetail.Year
	MonthyEMIs []EMI `json:"monthly_emis"`

	// Year is the year count of the financial year in scope from the tenure of the loan
	Year int `json:"year"`
}

// Data is the request body which carries the loan information
// for strategy recommendation
type Data struct {

	// AnnualInterestRate is the annual percentage interest levied for the loan borrowed
	AnnualInterestRate float32 `json:"annual_interest_rate"`

	// LoanAmount is the total amount borrowed from lender
	LoanAmount int `json:"loan_amount"`

	// LoanTenure is the number of years this loan is suppose to be settle in
	LoanTenure int `json:"loan_tenure"`
}

// EMI struct defines the EMI resource paid on a monthly basis
type EMI struct {

	// EMIAmount is the total EMI amount paid to the lender
	EMIAmount int `json:"emi_amount"`

	// Month is the month of the EMI in a financial year
	Month int `json:"month"`

	// TowardsInterest is the amount from the EMIAmount paid towards interest of the loan
	TowardsInterest int `json:"towards_interest"`

	// Towards Principle is the amount from the EMIAmount paid towards principle of the loan
	TowardsPrinciple int `json:"towards_principle"`
}

// StrategyInfo defines the complete comparison of different strategies to settle loan
// in a wise way
type StrategyInfo struct {

	// AnnualDetails is the array of AnnualDetail having yearly interesting information
	// for StrategyInfo.TotalYearsToSettle years
	AnnualDetails []AnnualDetail `json:"annual_details"`

	// AnnualEMIPercentageHike is the percentage hike on the existing EMI amount
	AnnualEMIPercentageHike int `json:"annual_emi_percentage_hike"`

	// TotalInterest is the total interest paid for the loan in StrategyInfo.TotalYearsToSettle
	TotalInterest int `json:"total_interest"`

	// TotalTaxSaved is the total taxes saved in StrategyInfo.TotalYearsToSettle
	TotalTaxsaved int `json:"total_tax_saved"`

	// TotalYearsToSettle is the total number of years it would take to settle the entire loan
	// amount with the StrategyInfo.AnnualEMIPercentageHike hike in EMIs
	TotalYearsToSettle int `json:"total_years_to_settle"`
}
