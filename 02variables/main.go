package main

import "fmt"

const LoginToken = "hahaha" // As the first letter (L) is capital, which means this is now public and can be accessed by any other file

func main() {
	var username string = "Satya"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type %T \n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T \n", smallVal)

	var smallFloat float32 = 255.8549385827453845803
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T \n", smallFloat)

	// default values and some aliases

	var anothervar int
	fmt.Println(anothervar)
	fmt.Printf("Value of another var: %T \n ", anothervar)

	// implicit type

	var website = "SatyasServer.com"
	fmt.Println(website)

	// no var style
	noOfUsers := 300000 // walrus operator : used for type reference so that we don't have to declare the type of a variable

	fmt.Println(noOfUsers) // This type of declaration is allowed only inside a function

	fmt.Println(LoginToken)
	fmt.Printf("The type of the variable is %T", LoginToken)

}
