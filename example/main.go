package main

import (
	"fmt"
	"math/rand"

	"github.com/jhunt/go-sample"
)

func main() {
	r := sample.NewReservoir(100)
	for i := 0; i < 2000; i++ {
		r.Sample(rand.Float64())
	}
	fmt.Printf("minimum is %v\n", r.Minimum())
	fmt.Printf("maximum is %v\n", r.Maximum())
	fmt.Printf(" median is %v\n", r.Median())
	fmt.Printf("st. dev is %v\n", r.Stdev())
}
