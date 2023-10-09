package improved

import "math"

// DCF calculates the intrinsic value of an investment based on discounted cash flows
// cr represents current earnings
// gr represents growth rate of earnings in 10 years. gr must be normalized
// fgr represents growth rate of earings more than 10 years later
// dr represents discounted rate
func DCF(cr, gr, fgr, dr float64, y int) float64 {
	v := 0.0 // Intrinsic value
	for i := 1; i <= 10; i++ {
		cf := cr * pow((1+gr), float64(i)) // Cash flow in the year
		dv := cf / pow((1+dr), float64(i)) // Discounted value of the year's cash flow
		v += dv                            // Add current year's discounted cash flow to the intrinsic value
	}
	for j := 11; j <= y; j++ {
		cf2 := cr * pow((1+gr), 10.0) * pow((1+fgr), float64(j-10)) // Cash Flow in the year for 10+ years
		dv2 := cf2 / pow((1+dr), float64(j))                        // Discounted value of the year's cash flow for 10+ years
		v += dv2
	}
	return math.Round(v)
}

// pow calculates the power of a float64 base to an exponent
func pow(b, e float64) float64 {
	if e == 0.0 {
		return 1
	}
	r := b // result
	for i := 1.0; i < e; i++ {
		r *= b
	}
	return r
}
