package main

import (
	"fmt"
	"hello-build-with-golang/wadd/calc"
	"hello-build-with-golang/wadd/util"
)

func main() {
	num1 := util.RandomInt(1, 1000)
	num2 := util.RandomInt(1, 1000)

	num3 := calc.Add(num1, num2)
	fmt.Println("result is ", num3)
}
