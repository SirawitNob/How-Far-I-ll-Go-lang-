package main

import (
	"fmt"
	"os"
)

func main(){
	writeFile();
	readFile();
}

func readFile(){
	fileName := "test.txt";
	fi, err := os.ReadFile(fileName);
	if err != nil {
		panic(err);
	}
	fmt.Printf("%s", fi);
}

func writeFile(){
	fileName := "test.txt";
	file, err := os.Create(fileName);
	if err != nil {
		panic(err);
	}
	length, err := file.WriteString("Hello World");
	if err != nil {
		panic(err);
	}
	fmt.Println("Write content on file : ", file.Name());
	fmt.Printf("file length : %d", length);
}