package main

import "fmt"

func main(){
	a := 25;
	fmt.Println(a);

	b, c, d, _, f := 1, 2, 3, 4, "happiness";
	fmt.Println(b, c, d, f);

	var g int;
	fmt.Println(g);
	g = 44;
	fmt.Println(g);

	var h string;
	h = "Hello World";
	fmt.Println(h);

	adam := 42;
	fmt.Printf("42 as a binary, %b\n", adam);
	fmt.Printf("42 as a hexadecimal, %x\n", adam);

	fmt.Printf("%v \t %b \t %x\n", b, b, b);
	fmt.Printf("%v \t %b \t %x\n", c, c, c);
	fmt.Printf("%v \t %b \t %x\n", d, d, d);
}