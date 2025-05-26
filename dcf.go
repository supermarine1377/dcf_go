package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/supermarine1377/dcf_go/src/dcf"
	"github.com/supermarine1377/dcf_go/src/dcf/condition"
)

type Input struct {
	c *condition.Condition
}

// NewInput creates a new Input instance by prompting the user for input.
func NewInput() (*Input, error) {
	ce, err := promptAndReadFloat(msgCurrentEarnings)
	if err != nil {
		return nil, err
	}
	gr, err := promptAndReadFloat(msgGrowthRate)
	if err != nil {
		return nil, err
	}
	tgr, err := promptAndReadFloat(msgTerminalGrowthRate)
	if err != nil {
		return nil, err
	}
	d, err := promptAndReadFloat(msgDiscountRate)
	if err != nil {
		return nil, err
	}
	y, err := promptAndReadFloat(msgYears)
	if err != nil {
		return nil, err
	}
	c, err := condition.New(
		condition.WithCurrentEarnings(ce),
		condition.WithGrowthRate(gr),
		condition.WithTerminalGrowthRate(tgr),
		condition.WithDiscountRate(d),
		condition.WithYears(int(y)),
	)
	if err != nil {
		return nil, err
	}

	return &Input{c: c}, nil
}

// promptAndReadFloat prompts the user for input and reads a float64 value from stdin.
func promptAndReadFloat(promptText string) (float64, error) {
	fmt.Print(promptText)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("failed to read input for '%s': %w", strings.TrimSpace(promptText), err)
	}
	input = strings.TrimSpace(input)

	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse '%s' as float for '%s': %w", input, strings.TrimRight(promptText, ": "), err)
	}
	return val, nil
}

// DCF calculates the discounted cash flow based on the input parameters.
func (i *Input) DCF() float64 {
	return dcf.DCF(i.c)
}
