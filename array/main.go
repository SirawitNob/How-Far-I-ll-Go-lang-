package main

import (
	"fmt"
	"time"
)

func main(){

	fmt.Println("Hello World");

	arr:=[4]string{"geek", "gfg", "NobNab", "Sirar"};
	arr2:=[...]int{1,2,3,4,5,6};
	sum := 0;

	for i:=0 ; i<len(arr) ; i++{
		fmt.Println(arr[i]);
	}

	for i:=0 ; i<len(arr2) ; i++{
		switch i {
		case 1:
			fmt.Println("Number 1.");
		case 2:
			fmt.Println("Number 2.");
		default:
			fmt.Println("Not 1 and 2.");
		}
		sum += arr2[i];
	}

	t := time.Now();
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon.");
	default:
		fmt.Println("It's after noon.");
	}

	fmt.Println("Sum of arr2 :", sum);

	for_range_loop(arr);
}

func for_range_loop(arr [4]string){
	for i, v := range arr {
		fmt.Printf("index : %d, \nvalue : %s\n", i, v);
	}
}