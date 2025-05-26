package condition

// Condition represents the parameters needed to calculate the intrinsic value of a stock using the Discounted Cash Flow (DCF) method.
type Condition struct {
	// currentEarnings represents current earnings
	currentEarnings float64
	// growthRate represents growth rate of earnings in 10 years. gr must be normalized
	growthRate float64
	// terminalGrowthRate represents termial growth rate of earnings more than 10 years later.
	// terminalGrowthRate must be normalized
	terminalGrowthRate float64
	// discountRate represents discounted rate
	discountRate float64
	// years represents the number of years to calculate
	years int
}

// Options are functions that modify a Condition instance.
type Option func(*Condition)

// New creates a new Condition instance with the provided parameters.
func New(options ...Option) *Condition {
	c := &Condition{}

	for _, option := range options {
		option(c)
	}
	return c
}

// WithCurrentEarnings sets the CurrentEarnings field.
func WithCurrentEarnings(cr float64) Option {
	return func(c *Condition) {
		c.currentEarnings = cr
	}
}

// WithGrowthRate sets the GrowthRate field.
func WithGrowthRate(gr float64) Option {
	return func(c *Condition) {
		c.growthRate = gr
	}
}

// WithTerminalGrowthRate sets the TerminalGrowthRate field.
func WithTerminalGrowthRate(tgr float64) Option {
	return func(c *Condition) {
		c.terminalGrowthRate = tgr
	}
}

// WithDiscountRate sets the DiscountRate field.
func WithDiscountRate(dr float64) Option {
	return func(c *Condition) {
		c.discountRate = dr
	}
}

// WithYears sets the Years field.
func WithYears(y int) Option {
	return func(c *Condition) {
		c.years = y
	}
}

// Cr returns the current earnings.
func (c *Condition) CurrentEarnings() float64 {
	return c.currentEarnings
}

// Gr returns the growth rate of earnings in 10 years.
func (c *Condition) GrowthRate() float64 {
	return c.growthRate
}

// Tgr returns the growth rate of earnings more than 10 years later.
func (c *Condition) TeminalGrowthRate() float64 {
	return c.terminalGrowthRate
}

// Dr returns the discounted rate.
func (c *Condition) DiscountRate() float64 {
	return c.discountRate
}

// Years returns the number of years to calculate.
func (c *Condition) Years() int {
	return c.years
}
