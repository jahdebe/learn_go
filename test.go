package main

import (
    "fmt"
    "strconv"
)

func Convert(input int) string {
	output := ""
	if input%3 == 0	{
		output += "Pling"
	}
	if input%5 == 0	{
		output += "Plang"
	}
	if input%7 == 0	{
		output += "Plong"
	}
	if (input%3 != 0 && input%5 != 0 && input%7 != 0)	{
		output = strconv.Itoa(input)
		// output = string(input)
	}
	return output  
}

func main() {
	fmt.Println("Hello, world.")
	test := ""
	fmt.Println(test)
	test += "more"
	fmt.Println(test)
	test += "MOAR"
	fmt.Println(test)
}