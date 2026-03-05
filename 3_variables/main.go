package main

import "fmt"

func main() {

	//Declare a variable -> var variable_name variable_type
	var name string = "go lang "

	// Whenever we create a variable in GoLang we have to use it. If we arent using it need to delete.
	fmt.Println(name)

	// Sometimes if we dont define var type , Go Lang is able to infer the var type
	var surname = "Devansh"

	fmt.Printf("Type of variable : %T \n", surname)

	fmt.Printf("This is the value : %s", surname)

	fmt.Println("This is the value : ", surname)

	//shorthand syntax
	first_name := "golang"

	fmt.Println("This is the shorthand version:", first_name)

	// Question arises what is even the need of var or defining the type when we can just use shorthand

	// Var still needs to be used when we are declaring variables

	var declare_name string

	declare_name = "Test Demo Go Lang"

	fmt.Println("This is the declared variable : ", declare_name)

}
