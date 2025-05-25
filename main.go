package main

import (
	"fmt"

	"github.com/supermarine1377/dcf_go/src"
)

func main() {
	i, err := src.NewInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("intrinsic value: %v\n", i.DCF())
}
