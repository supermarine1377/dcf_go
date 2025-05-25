package src

import (
	"fmt"
	"strconv"

	"github.com/supermarine1377/dcf_go/src/lib/dcf"
	"github.com/supermarine1377/dcf_go/src/lib/dcf/condition"
)

type Input struct {
	currentEarnings  float64
	growthRate       float64
	futureGrowthRate float64
	discountRate     float64
	years            int
}

func NewInput() (*Input, error) {
	fmt.Printf("Current earings: ")
	ce, err := fromStdin()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Growth rate of earnings in 10 years (must be normalized): ")
	gr, err := fromStdin()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Growth rate of earings more than 10 years later (must be normalized): ")
	fgr, err := fromStdin()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Discount rate (must be normalized): ")
	d, err := fromStdin()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Year (must be greater than 10): ")
	y, err := fromStdin()
	if err != nil {
		return nil, err
	}
	return &Input{
		currentEarnings:  ce,
		growthRate:       gr,
		futureGrowthRate: fgr,
		discountRate:     d,
		years:            int(y),
	}, nil
}

func fromStdin() (float64, error) {
	var s string
	if _, err := fmt.Scan(&s); err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

func (i *Input) DCF() float64 {
	c := condition.New(
		i.currentEarnings,
		i.growthRate,
		i.futureGrowthRate,
		i.discountRate,
		i.years,
	)
	return dcf.DCF(c)
}
