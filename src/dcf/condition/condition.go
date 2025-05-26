package condition

import "fmt"

// Condition represents the parameters needed to calculate the intrinsic value of a stock using the Discounted Cash Flow (DCF) method.
type Condition struct {
	// currentEarnings represents current earnings
	currentEarnings float64
	// growthRate represents growth rate of earnings in 10 years. gr must be normalized
	growthRate float64
	// growthYears represents the number of years where the growth rate is applied.
	growthYears int
	// terminalGrowthRate represents termial growth rate of earnings more than 10 years later.
	// terminalGrowthRate must be normalized
	terminalGrowthRate float64
	// discountRate represents discounted rate
	discountRate float64
	// years represents the number of years to calculate. Must be greater than grrowthYears
	years int
}

// Options are functions that modify a Condition instance.
type Option func(*Condition)

// New creates a new Condition instance with the provided parameters.
func New(options ...Option) (*Condition, error) {
	c := &Condition{}

	for _, option := range options {
		option(c)
	}

	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("invalid condition: %w", err)
	}

	return c, nil
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

// WithGrowthYears sets the GrowthYears field.
func WithGrowthYears(gy int) Option {
	return func(c *Condition) {
		c.growthYears = gy
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

// Gr returns the growth rate of earnings in GrowthYears.
func (c *Condition) GrowthRate() float64 {
	return c.growthRate
}

// GrowthYears returns the number of years where the growth rate is applied.
func (c *Condition) GrowthYears() int {
	return c.growthYears
}

// Tgr returns the growth rate of earnings more than GrowthYears later.
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

// Validate validates the condition.
func (c *Condition) Validate() error {
	if c.currentEarnings <= 0 {
		return fmt.Errorf("current earnings must be greater than 0, got %f", c.currentEarnings)
	}
	if c.growthRate < 0 || c.growthRate > 1 {
		return fmt.Errorf("growth rate must be between 0 and 1, got %f", c.growthRate)
	}
	if c.terminalGrowthRate < 0 || c.terminalGrowthRate > 1 {
		return fmt.Errorf("terminal growth rate must be between 0 and 1, got %f", c.terminalGrowthRate)
	}
	if c.discountRate <= 0 || c.discountRate > 1 {
		return fmt.Errorf("discount rate must be between 0 and 1, got %f", c.discountRate)
	}
	if c.growthYears <= 0 || c.growthYears > c.years {
		return fmt.Errorf("growth years must be greater than 0 and less than or equal to years, got %d", c.growthYears)
	}

	if c.years <= 0 {
		return fmt.Errorf("years must be greater than 0, got %d", c.years)
	}
	return nil
}
