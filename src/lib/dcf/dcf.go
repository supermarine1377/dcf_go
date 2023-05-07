package dcf

import "math"

type cashFlow []float64

func DCF(currentCashFlow float64, growthRate float64, disountRate float64, evebit float64, years int) float64 {
	var futureCh cashFlow
	prev := currentCashFlow
	for y := 1; y < years+1; y++ {
		prev = prev * (1 + growthRate)
		futureCh = append(futureCh, prev)
	}
	var discounted []float64
	for i, fch := range futureCh {
		discounted = append(discounted, fch/math.Pow((1+disountRate), float64(i+1)))
	}
	var result float64
	for i := range discounted {
		if i == len(discounted)-1 {
			result += discounted[i] * evebit
		} else {
			result += discounted[i]
		}
	}
	return result
}
