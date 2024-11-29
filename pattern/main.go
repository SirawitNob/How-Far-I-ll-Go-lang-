package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello World");
	diamond(7);
}

func pyramid_half_right(hieght int){

	for i := 1 ; i <= hieght ; i++ {
		for j := 1 ; j <= i ; j++ {
			fmt.Printf("* ");
		}
		fmt.Println();
	}
}

func pyramid_half_left(height int){
	for i := height ; i > 0 ; i-- {
		for j := 1 ; j < i ; j++ {
			fmt.Printf("  ");
		}
		for k := i ; k <= height ; k++ {
			fmt.Printf("* ");
		}
		fmt.Println();
	}
}

func pyramid_full(height int){
	for i := height ; i > 0 ; i-- {
		for j := 0 ; j < i ; j++ {
			fmt.Printf(" ");
		}
		for k := i ; k <= height ; k++ {
			fmt.Printf("* ");
		}
		fmt.Println();
	}
}

func pyramid_full_reverse(height int){
	for i := height ; i > 0 ; i-- {
		for j := i ; j < height ; j++ {
			fmt.Printf(" ");
		}
		for k := i ; k > 0 ; k-- {
			fmt.Printf("* ");
		}
		fmt.Println();
	}
}

func diamond(height int){
	var first_half, second_half int;
	if height % 2 != 0 {
		first_half = (height / 2) + 1;
	} else {
		first_half = (height / 2);
	}
	second_half = height - first_half;
	for i := 0 ; i < first_half ; i++ {
		for j := first_half - i ; j > 0 ; j-- {
			fmt.Printf(" ");
		}
		for k := 0 ; k <= i ; k++ {
			fmt.Printf("* ");
		}
		fmt.Println();
	}
	for i := second_half ; i > 0 ; i-- {
		for j := second_half - i ; j >= 0 ; j-- {
			fmt.Printf(" ");
		}
		for k := 0 ; k < i ; k++ {
			fmt.Printf(" *");
		}
		fmt.Println();
	}
}