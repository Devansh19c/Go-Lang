package main

import "fmt"

const age int = 30

var name string = "golang"

/*

name := "golang"

The := operator performs short variable declaration and is only valid inside functions because it relies on local scope.
At the package level Go requires explicit declarations using var, const, type, or func.

*/

func main() {

	const name = "GoLang"

	// 	const name := "GoLang". Shorthand will not work because it is only meant for var.  So := always means: Create a variable and infer its type IMP - Only inside a Function Scope

	// name = "Trying to change".   This will not let us change value of name.

	fmt.Println(age)
	fmt.Printf("%T", age)
}
