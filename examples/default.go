package main

import (
	"fmt"

	. "github.com/kissgyorgy/missingo"
)

func main() {
	defaultPrint()
	fmt.Println()
	defaultPrint("value given")
}

func defaultPrint(values ...string) {
	result := Default("default value").IfNotSet(values)
	fmt.Printf("Result: %v, %T\n", result, result)
}
