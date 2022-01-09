package homeloan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

var taxSlab = 0.3

func computeStandardEMI(data Data) float64 {
	var (
		loanAmount          = data.LoanAmount
		monthlyInterestRate = data.AnnualInterestRate / 1200
		loanTenure          = data.LoanTenure * 12
	)

	factor := math.Pow((1 + monthlyInterestRate), loanTenure)
	standardEMI := loanAmount * monthlyInterestRate * factor / (factor - 1)
	return standardEMI
}

func updateAnnualDetails(annualDetail AnnualDetail, strategyInfo *StrategyInfo) {

	for _, emi := range annualDetail.MonthyEMIs {
		annualDetail.AnnualInterest += emi.TowardsInterest
		annualDetail.AnnualEMIAmount += emi.EMIAmount
	}

	annualDetail.AnnualTaxSaved = annualDetail.AnnualInterest * taxSlab
	strategyInfo.AnnualDetails = append(strategyInfo.AnnualDetails, annualDetail)
}

func updateStrategyInfo(strategyInfo *StrategyInfo) {
	for _, info := range strategyInfo.AnnualDetails {
		strategyInfo.TotalInterest += int(info.AnnualInterest)
		strategyInfo.TotalTaxsaved += int(info.AnnualTaxSaved)
	}
	strategyInfo.TotalYearsToSettle = len(strategyInfo.AnnualDetails)
}

func computeStrategies(data Data) []StrategyInfo {

	var (
		recommendations []StrategyInfo
	)

	for yearlyPercentageHike := 0; yearlyPercentageHike <= 20; yearlyPercentageHike += 5 {
		var (
			principle    = data.LoanAmount
			annualDetail AnnualDetail
			strategyInfo StrategyInfo
			standardEMI  = computeStandardEMI(data)
		)

		for month := 1; month < int(data.LoanTenure*12)+1; month++ {
			if principle <= 0 {
				break
			}

			var emi EMI

			emi.Month = month
			emi.EMIAmount = standardEMI
			emi.TowardsInterest = principle * data.AnnualInterestRate / 1200
			emi.TowardsPrinciple = standardEMI - emi.TowardsInterest

			principle = principle - emi.TowardsPrinciple

			annualDetail.MonthyEMIs = append(annualDetail.MonthyEMIs, emi)

			if month%12 == 0 {
				annualDetail.Year = month / 12
				standardEMI = standardEMI * float64(1+float64(yearlyPercentageHike)/100.0)
				updateAnnualDetails(annualDetail, &strategyInfo)
				annualDetail = AnnualDetail{}
			}
		}
		strategyInfo.AnnualEMIPercentageHike = float64(yearlyPercentageHike)
		updateStrategyInfo(&strategyInfo)
		recommendations = append(recommendations, strategyInfo)
	}
	return recommendations
}

// Home is route handler than returns the home page
func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// Strategies function accepts HTTP POST requests to compute
// recommendation for optimum home loan settlement strategies
func Strategies(w http.ResponseWriter, r *http.Request) {

	var data *Data
	var requestBody []byte

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error while reading request body %s", err.Error())
	}

	err = json.Unmarshal(requestBody, &data)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Error while parsing the request body: %s", err.Error())
	}

	recommendations := computeStrategies(*data)
	json.NewEncoder(w).Encode(recommendations)
}
