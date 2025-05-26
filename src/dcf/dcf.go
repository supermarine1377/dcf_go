package dcf

import (
	"math"

	"github.com/supermarine1377/dcf_go/src/dcf/condition"
)

func DCF(c *condition.Condition) float64 {
	var (
		cr  = c.CurrentEarnings()
		gr  = c.GrowthRate()
		tgr = c.TeminalGrowthRate()
		dr  = c.DiscountRate()
		y   = c.Years()
	)

	// Calculate the discounted cash flow for the first 10 years
	presentValue := calculatePresentValue(cr, gr, dr, 10)

	// Calculate the discounted cash flow for the remaining years
	presentValue += calculatePresentValueLongTerm(cr, gr, tgr, dr, y)

	return math.Round(presentValue)
}

// calculatePresentValue calculates the discounted cash flow for the first 10 years
func calculatePresentValue(currentEarnings, growthRate, discountRate float64, years int) float64 {
	presentValue := 0.0
	for i := 1; i <= years; i++ {
		cashFlow := currentEarnings * math.Pow((1+growthRate), float64(i))
		discountedValue := cashFlow / math.Pow((1+discountRate), float64(i))
		presentValue += discountedValue
	}
	return presentValue
}

// calculatePresentValueLongTerm calculates the discounted cash flow for years after the first 10 years
func calculatePresentValueLongTerm(currentEarnings, growthRate, terminalGrowthRate, discountRate float64, years int) float64 {
	presentValue := 0.0
	for j := 11; j <= years; j++ {
		// Cash Flow in the year for 10+ years
		cf2 := currentEarnings * math.Pow((1+growthRate), 10) * math.Pow((1+terminalGrowthRate), float64(j-10))
		// Discounted value of the year's cash flow for 10+ years
		dv2 := cf2 / math.Pow((1+discountRate), float64(j))
		presentValue += dv2
	}
	return presentValue
}
