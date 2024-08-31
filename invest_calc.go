package main

import (
	"fmt"
	"math"
)

func main() {
	var investAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futurevalue = float64(investAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println(futurevalue)
}
