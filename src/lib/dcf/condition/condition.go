package condition

type Condition struct {
	// cr represents current earnings
	cr float64
	// gr represents growth rate of earnings in 10 years. gr must be normalized
	gr float64
	// fgr represents growth rate of earings more than 10 years later
	fgr float64
	// dr represents discounted rate
	dr float64
	// y represents the number of years to calculate
	y int
}

// New creates a new Condition instance with the provided parameters.
func New(cr, gr, fgr, dr float64, y int) *Condition {
	return &Condition{
		cr:  cr,
		gr:  gr,
		fgr: fgr,
		dr:  dr,
		y:   y,
	}
}

// Cr returns the current earnings.
func (c *Condition) CurrentEarnings() float64 {
	return c.cr
}

// Gr returns the growth rate of earnings in 10 years.
func (c *Condition) GrowthRate() float64 {
	return c.gr
}

// Fgr returns the growth rate of earnings more than 10 years later.
func (c *Condition) FutureGrowthRate() float64 {
	return c.fgr
}

// Dr returns the discounted rate.
func (c *Condition) DiscountRate() float64 {
	return c.dr
}

// Years returns the number of years to calculate.
func (c *Condition) Years() int {
	return c.y
}
