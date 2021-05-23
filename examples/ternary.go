package main

import (
	"fmt"

	. "github.com/kissgyorgy/missingo"
)

func main() {
	trueValue := If(true).Then("true value").Else("false value")
	fmt.Printf("Result: %q, %T\n", trueValue, trueValue)

	falseValue := If(true).Then("true value").Else("false value")
	fmt.Printf("Result: %q, %T\n", falseValue, falseValue)
}
