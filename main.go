package main

import (
	"fmt"
	"github.com/abadojack/whatlanggo"
	"github.com/jinzhu/now"
	"github.com/montanaflynn/stats"
)

func main() {
	info := whatlanggo.Detect("Hola que tal estas?")
	fmt.Println("Language:", whatlanggo.LangToString(info.Lang), "Script:", whatlanggo.Scripts[info.Script])
	fmt.Printf("Now is %v\n", now.BeginningOfMinute())

	var data = []float64{1, 2, 3, 4, 4, 5}

	median, _ := stats.Median(data)
	fmt.Println(median) // 3.5
}
