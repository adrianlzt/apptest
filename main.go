package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"github.com/montanaflynn/stats"
)

func main() {
	fmt.Printf("Now is %v\n", now.BeginningOfMinute())

	var data = []float64{1, 2, 3, 4, 4, 5}

	median, _ := stats.Median(data)
	fmt.Println(median) // 3.5
}
