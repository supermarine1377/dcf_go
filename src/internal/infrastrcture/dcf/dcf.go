package dcf

import (
	"fmt"
	"strconv"

	"github.com/supermarine1377/dcf_go/src/lib/dcf"
)

type Input struct {
	current        float64
	grothRate      float64
	discountedRate float64
	evebit         float64
	years          int
}

func NewInput() (*Input, error) {
	fmt.Print("Current earings: ")
	c, err := fromStdin()
	if err != nil {
		return nil, err
	}
	fmt.Print("Groth rate: ")
	g, err := fromStdin()
	if err != nil {
		return nil, err
	}
	fmt.Print("Discounted rate: ")
	d, err := fromStdin()
	if err != nil {
		return nil, err
	}
	fmt.Print("eveit: ")
	e, err := fromStdin()
	if err != nil {
		return nil, err
	}
	fmt.Print("year: ")
	y, err := fromStdin()
	if err != nil {
		return nil, err
	}
	return &Input{
		current:        c,
		grothRate:      g,
		discountedRate: d,
		evebit:         e,
		years:          int(y),
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
	return dcf.DCF(i.current, i.grothRate, i.discountedRate, i.evebit, i.years)
}
