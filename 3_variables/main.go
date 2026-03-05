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

}
