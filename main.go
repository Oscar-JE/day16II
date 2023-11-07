package main

import (
	"day16/parse"
	"fmt"
)

func main() {
	fileName := "input_short.txt"
	inputs := parse.ParseFile(fileName)
	for _, in := range inputs {
		fmt.Println(in)
	}
}
