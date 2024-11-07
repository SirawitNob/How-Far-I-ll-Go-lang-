package main

import "fmt"

func main(){
	var str string = sayHello();
	var rtn int;
	sayHello();
	rtn = addNumber(15, 25);
	fmt.Println(str);
	fmt.Println(rtn);
}

func addNumber(x int, y int) int {
	var ans int = 40;
	ans += x + y;
	return ans;
}

func sayHello() string{
	return "Hello World";
}