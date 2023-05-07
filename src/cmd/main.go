package main

import (
	"fmt"

	"github.com/supermarine1377/dcf_go/src/internal/infrastrcture/dcf"
)

func main() {
	i, err := dcf.NewInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("intrinsic value: %v\n", i.DCF())
}
