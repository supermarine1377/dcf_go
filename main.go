package main

import (
	"fmt"
)

func main() {
	i, err := NewInput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("intrinsic value: %v\n", i.DCF())
}
