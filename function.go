package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(PI())
	fmt.Println((hello("おまモア")))
	fmt.Println(hello2("fdjaksl", "おまおま"))
	fmt.Println(multipleArgs("fsa", "ioimh"))
}

func PI() float64 {
	return math.Pi
}

func hello(arg string) string {
	return arg
}

func hello2(arg1, arg2 string) string {
	return arg1 + " " + arg2
}

func multipleArgs(arg1, arg2 string) (string, string) {
	return arg1, arg2
}
