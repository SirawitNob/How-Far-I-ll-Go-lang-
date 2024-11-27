package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(time.Now());
	fmt.Println(math.Min(15, 2));
	var number int = -5;
	if number < 0 {
		fmt.Println(absolute_number(number));
	}
	
	fmt.Println(math.Pi);
}

func absolute_number(number int) int {
	return int(math.Abs(float64(number)));
}