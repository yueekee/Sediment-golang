package main

import "fmt"

func callFunction(passedFunction func()) {
	passedFunction()
}
func callTwice(passedFunction func()) {
	passedFunction()
	passedFunction()
}
func callWithArgument(passedFunction func(string,bool)) {
	passedFunction("This sentence is", false)
}
func printReturnValue(passedFunction func() string) {
	fmt.Println(passedFunction())
}

func functionA() {
	fmt.Println("function called")
}
func functionB() string {
	fmt.Println("function called")
	return "Return from function"
}
func functionC(a string, b bool) {
	fmt.Println("function called")
	fmt.Println(a, b)
}

func main() {
	callFunction(functionA)
	callTwice(functionA)
	callWithArgument(functionC)
	printReturnValue(functionB)
}

/*
function called
function called
function called
function called
This sentence is false
function called
Return from function
*/